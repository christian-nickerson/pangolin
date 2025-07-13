use crate::distance::{cosine::cosine_similarity, euclidean::l2_distance, error::DistanceError};
use crate::hnsw::graph::{DistanceMetric, EdgeWeight, VectorGraph, VectorNode};
use std::collections::{BinaryHeap, HashSet};
use std::cmp::Ordering;

/// Configuration parameters for HNSW index construction
#[derive(Debug, Clone)]
pub struct HnswConfig {
    /// Maximum number of connections for each node at layer 0
    pub m: usize,
    /// Maximum number of connections for each node at layers > 0
    pub max_m: usize,
    /// Maximum number of connections during construction
    pub max_m_l: usize,
    /// Level generation factor (typically 1/ln(2))
    pub ml: f64,
    /// Number of candidates to consider during search
    pub ef_construction: usize,
    /// Distance metric to use
    pub metric: DistanceMetric,
}

impl Default for HnswConfig {
    fn default() -> Self {
        Self {
            m: 16,
            max_m: 16,
            max_m_l: 16,
            ml: 1.0 / (2.0_f64).ln(),
            ef_construction: 200,
            metric: DistanceMetric::Euclidean,
        }
    }
}

/// A candidate node with its distance for priority queue operations
#[derive(Debug, Clone)]
struct Candidate {
    node_id: u64,
    distance: f32,
}

impl PartialEq for Candidate {
    fn eq(&self, other: &Self) -> bool {
        self.distance == other.distance
    }
}

impl Eq for Candidate {}

impl PartialOrd for Candidate {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Candidate {
    fn cmp(&self, other: &Self) -> Ordering {
        other.distance.partial_cmp(&self.distance).unwrap_or(Ordering::Equal)
    }
}

/// HNSW Index builder and search structure
pub struct HnswIndex {
    graph: VectorGraph,
    config: HnswConfig,
    entry_point: Option<u64>,
    rng: fastrand::Rng,
}

impl HnswIndex {
    /// Creates a new HNSW index with the given configuration
    pub fn new(config: HnswConfig) -> Self {
        let graph = VectorGraph::new(config.metric, config.m);
        Self {
            graph,
            config,
            entry_point: None,
            rng: fastrand::Rng::new(),
        }
    }

    /// Creates a new HNSW index with default configuration
    pub fn with_default_config() -> Self {
        Self::new(HnswConfig::default())
    }

    /// Returns the number of nodes in the index
    pub fn len(&self) -> usize {
        self.graph.node_count()
    }

    /// Returns true if the index is empty
    pub fn is_empty(&self) -> bool {
        self.graph.node_count() == 0
    }

    /// Calculates distance between two vectors based on the configured metric
    fn calculate_distance(&self, vec1: &[f32], vec2: &[f32]) -> Result<f32, DistanceError> {
        let vec1_f64: Vec<f64> = vec1.iter().map(|&x| x as f64).collect();
        let vec2_f64: Vec<f64> = vec2.iter().map(|&x| x as f64).collect();

        match self.config.metric {
            DistanceMetric::Euclidean => l2_distance(&vec1_f64, &vec2_f64).map(|d| d as f32),
            DistanceMetric::Cosine => {
                cosine_similarity(&vec1_f64, &vec2_f64).map(|sim| (1.0 - sim) as f32)
            }
        }
    }

    /// Generates a random layer for a new node
    fn get_random_layer(&mut self) -> usize {
        let mut layer = 0;
        while self.rng.f64() < 0.5 && layer < 16 {
            layer += 1;
        }
        layer
    }

    /// Searches for the closest nodes at a given layer
    fn search_layer(
        &self,
        query: &[f32],
        entry_points: Vec<u64>,
        num_closest: usize,
        layer: usize,
    ) -> Result<Vec<Candidate>, DistanceError> {
        let mut visited = HashSet::new();
        let mut candidates = BinaryHeap::new();
        let mut w = BinaryHeap::new();

        for ep in entry_points {
            if let Some(node) = self.graph.get_node(ep) {
                if node.layer >= layer {
                    let distance = self.calculate_distance(query, &node.vector)?;
                    let candidate = Candidate {
                        node_id: ep,
                        distance,
                    };
                    candidates.push(candidate.clone());
                    w.push(std::cmp::Reverse(candidate));
                    visited.insert(ep);
                }
            }
        }

        while let Some(candidate) = candidates.pop() {
            if let Some(std::cmp::Reverse(furthest)) = w.peek() {
                if candidate.distance > furthest.distance {
                    break;
                }
            }

            if let Some(neighbors) = self.graph.get_neighbors(candidate.node_id) {
                for neighbor_id in neighbors {
                    if !visited.contains(&neighbor_id) {
                        visited.insert(neighbor_id);

                        if let Some(neighbor_node) = self.graph.get_node(neighbor_id) {
                            if neighbor_node.layer >= layer {
                                let distance = self.calculate_distance(query, &neighbor_node.vector)?;
                                let neighbor_candidate = Candidate {
                                    node_id: neighbor_id,
                                    distance,
                                };

                                if w.len() < num_closest {
                                    candidates.push(neighbor_candidate.clone());
                                    w.push(std::cmp::Reverse(neighbor_candidate));
                                } else if let Some(std::cmp::Reverse(furthest)) = w.peek() {
                                    if distance < furthest.distance {
                                        candidates.push(neighbor_candidate.clone());
                                        w.pop();
                                        w.push(std::cmp::Reverse(neighbor_candidate));
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }

        Ok(w.into_iter().map(|std::cmp::Reverse(c)| c).collect())
    }

    /// Selects M neighbors using a simple heuristic
    fn select_neighbors_simple(&self, candidates: Vec<Candidate>, m: usize) -> Vec<u64> {
        candidates
            .into_iter()
            .take(m)
            .map(|c| c.node_id)
            .collect()
    }

    /// Inserts a new vector into the HNSW index
    pub fn insert(&mut self, id: u64, vector: Vec<f32>) -> Result<(), DistanceError> {
        if self.graph.contains_node(id) {
            return Ok(()); // Node already exists
        }

        let layer = self.get_random_layer();
        let max_connections = if layer == 0 { self.config.m } else { self.config.max_m };

        let node = VectorNode::new(id, vector.clone(), layer, max_connections);
        self.graph.add_node(node);

        if self.entry_point.is_none() {
            self.entry_point = Some(id);
            return Ok(());
        }

        let entry_point = self.entry_point.unwrap();
        let max_layer = self.graph.max_layer().unwrap_or(0);

        // Search from top layer down to layer + 1
        let mut current_closest = vec![entry_point];
        for lc in (layer + 1..=max_layer).rev() {
            current_closest = self.search_layer(&vector, current_closest, 1, lc)?
                .into_iter()
                .map(|c| c.node_id)
                .collect();
        }

        // Search and connect from layer down to 0
        for lc in (0..=layer).rev() {
            let candidates = self.search_layer(&vector, current_closest.clone(), self.config.ef_construction, lc)?;

            let m = if lc == 0 { self.config.m } else { self.config.max_m };
            let selected_neighbors = self.select_neighbors_simple(candidates.clone(), m);

            // Add bidirectional connections
            for neighbor_id in &selected_neighbors {
                let distance = self.calculate_distance(&vector, &self.graph.get_node(*neighbor_id).unwrap().vector)?;
                let weight = EdgeWeight::new(distance, self.config.metric);
                self.graph.add_edge(id, *neighbor_id, weight);
            }

            // Prune connections of neighbors if needed
            for neighbor_id in &selected_neighbors {
                if let Some(neighbor_connections) = self.graph.get_neighbors(*neighbor_id) {
                    if neighbor_connections.len() > m {
                        // Simple pruning: keep closest connections
                        let neighbor_vector = &self.graph.get_node(*neighbor_id).unwrap().vector;
                        let mut neighbor_candidates = Vec::new();

                        for conn_id in neighbor_connections {
                            if conn_id != id {
                                let conn_vector = &self.graph.get_node(conn_id).unwrap().vector;
                                let distance = self.calculate_distance(neighbor_vector, conn_vector)?;
                                neighbor_candidates.push(Candidate {
                                    node_id: conn_id,
                                    distance,
                                });
                            }
                        }

                        neighbor_candidates.sort_by(|a, b| a.distance.partial_cmp(&b.distance).unwrap());
                        let to_keep: HashSet<u64> = neighbor_candidates
                            .into_iter()
                            .take(m - 1)
                            .map(|c| c.node_id)
                            .collect();
                        to_keep.insert(id);

                        // Remove excess connections
                        if let Some(all_neighbors) = self.graph.get_neighbors(*neighbor_id) {
                            for conn_id in all_neighbors {
                                if !to_keep.contains(&conn_id) {
                                    self.graph.remove_edge(*neighbor_id, conn_id);
                                }
                            }
                        }
                    }
                }
            }

            current_closest = candidates.into_iter().map(|c| c.node_id).collect();
        }

        // Update entry point if necessary
        if layer > max_layer {
            self.entry_point = Some(id);
        }

        Ok(())
    }

    /// Searches for the k nearest neighbors to the query vector
    pub fn search(&self, query: &[f32], k: usize, ef: usize) -> Result<Vec<(u64, f32)>, DistanceError> {
        if self.entry_point.is_none() {
            return Ok(Vec::new());
        }

        let entry_point = self.entry_point.unwrap();
        let max_layer = self.graph.max_layer().unwrap_or(0);

        // Search from top layer down to layer 1
        let mut current_closest = vec![entry_point];
        for lc in (1..=max_layer).rev() {
            current_closest = self.search_layer(query, current_closest, 1, lc)?
                .into_iter()
                .map(|c| c.node_id)
                .collect();
        }

        // Search at layer 0 with ef
        let candidates = self.search_layer(query, current_closest, ef.max(k), 0)?;

        Ok(candidates
            .into_iter()
            .take(k)
            .map(|c| (c.node_id, c.distance))
            .collect())
    }

    /// Builds the index from a collection of vectors
    pub fn build_from_vectors(&mut self, vectors: Vec<(u64, Vec<f32>)>) -> Result<(), DistanceError> {
        for (id, vector) in vectors {
            self.insert(id, vector)?;
        }
        Ok(())
    }

    /// Returns a reference to the underlying graph
    pub fn graph(&self) -> &VectorGraph {
        &self.graph
    }

    /// Returns the entry point of the index
    pub fn entry_point(&self) -> Option<u64> {
        self.entry_point
    }

    /// Returns the configuration used by this index
    pub fn config(&self) -> &HnswConfig {
        &self.config
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hnsw_basic_operations() {
        let mut index = HnswIndex::with_default_config();

        // Test insertion
        let vector1 = vec![1.0, 2.0, 3.0];
        let vector2 = vec![4.0, 5.0, 6.0];

        assert!(index.insert(1, vector1.clone()).is_ok());
        assert!(index.insert(2, vector2.clone()).is_ok());

        assert_eq!(index.len(), 2);
        assert!(!index.is_empty());

        // Test search
        let query = vec![1.5, 2.5, 3.5];
        let results = index.search(&query, 1, 10).unwrap();

        assert_eq!(results.len(), 1);
        assert_eq!(results[0].0, 1); // Should find vector1 as closest
    }

    #[test]
    fn test_hnsw_build_from_vectors() {
        let mut index = HnswIndex::with_default_config();

        let vectors = vec![
            (1, vec![1.0, 0.0]),
            (2, vec![0.0, 1.0]),
            (3, vec![-1.0, 0.0]),
            (4, vec![0.0, -1.0]),
        ];

        assert!(index.build_from_vectors(vectors).is_ok());
        assert_eq!(index.len(), 4);

        // Search for closest to (1, 0)
        let results = index.search(&[1.0, 0.0], 2, 10).unwrap();
        assert_eq!(results.len(), 2);
        assert_eq!(results[0].0, 1); // Exact match should be first
    }
}
