use petgraph::graph::{NodeIndex, UnGraph};
use petgraph::Graph;
use std::collections::HashMap;

/// Represents a vector node in the HNSW graph
#[derive(Debug, Clone)]
pub struct VectorNode {
    pub id: u64,
    pub vector: Vec<f32>,
    /// layer 0 is bottom
    pub layer: usize,
    pub max_connections: usize,
}

/// Represents the weight/distance of an edge between two nodes
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct EdgeWeight {
    pub distance: f32,
    pub metric: DistanceMetric,
}

/// Supported distance metrics for edges
#[derive(Debug, Clone, Copy, PartialEq)]
pub enum DistanceMetric {
    Cosine,
    Euclidean,
}

/// The main vector graph structure using petgraph's UnStableGraph
pub struct VectorGraph {
    graph: UnGraph<VectorNode, EdgeWeight>,
    /// Mapping from external node IDs to internal NodeIndex
    id_to_index: HashMap<u64, NodeIndex>,
    /// Mapping from internal NodeIndex to external node IDs
    index_to_id: HashMap<NodeIndex, u64>,
    default_metric: DistanceMetric,
    default_max_connections: usize,
}

impl VectorNode {
    /// Creates a new vector node
    pub fn new(id: u64, vector: Vec<f32>, layer: usize, max_connections: usize) -> Self {
        Self { id, vector, layer, max_connections }
    }

    /// Returns the dimensionality of the vector
    pub fn dimension(&self) -> usize {
        self.vector.len()
    }
}

impl EdgeWeight {
    /// Creates a new edge weight
    pub fn new(distance: f32, metric: DistanceMetric) -> Self {
        Self { distance, metric }
    }
}

impl VectorGraph {
    /// Creates a new empty vector graph
    pub fn new(default_metric: DistanceMetric, default_max_connections: usize) -> Self {
        Self {
            graph: Graph::new_undirected(),
            id_to_index: HashMap::new(),
            index_to_id: HashMap::new(),
            default_metric,
            default_max_connections,
        }
    }

    /// Returns the number of nodes in the graph
    pub fn node_count(&self) -> usize {
        self.graph.node_count()
    }

    /// Returns the number of edges in the graph
    pub fn edge_count(&self) -> usize {
        self.graph.edge_count()
    }

    /// Returns a reference to the underlying petgraph
    pub fn graph(&self) -> &UnGraph<VectorNode, EdgeWeight> {
        &self.graph
    }

    /// Returns a mutable reference to the underlying petgraph
    pub fn graph_mut(&mut self) -> &mut UnGraph<VectorNode, EdgeWeight> {
        &mut self.graph
    }

    /// Gets the NodeIndex for a given external ID
    pub fn get_node_index(&self, id: u64) -> Option<NodeIndex> {
        self.id_to_index.get(&id).copied()
    }

    /// Gets the external ID for a given NodeIndex
    pub fn get_node_id(&self, index: NodeIndex) -> Option<u64> {
        self.index_to_id.get(&index).copied()
    }

    /// Checks if a node with the given ID exists
    pub fn contains_node(&self, id: u64) -> bool {
        self.id_to_index.contains_key(&id)
    }

    /// Gets a reference to a node by its external ID
    pub fn get_node(&self, id: u64) -> Option<&VectorNode> {
        self.get_node_index(id)
            .and_then(|index| self.graph.node_weight(index))
    }

    /// Gets a mutable reference to a node by its external ID
    pub fn get_node_mut(&mut self, id: u64) -> Option<&mut VectorNode> {
        self.get_node_index(id)
            .and_then(move |index| self.graph.node_weight_mut(index))
    }

    /// Gets all neighbors of a node by its external ID
    pub fn get_neighbors(&self, id: u64) -> Option<Vec<u64>> {
        self.get_node_index(id).map(|index| {
            self.graph
                .neighbors(index)
                .filter_map(|neighbor_index| self.get_node_id(neighbor_index))
                .collect()
        })
    }

    /// Gets all neighbors with their edge weights for a node by its external ID
    pub fn get_neighbors_with_weights(&self, id: u64) -> Option<Vec<(u64, EdgeWeight)>> {
        self.get_node_index(id).map(|index| {
            self.graph
                .edges(index)
                .filter_map(|edge| {
                    let neighbor_id = self.get_node_id(edge.target())?;
                    Some((neighbor_id, *edge.weight()))
                })
                .collect()
        })
    }

    /// Gets the edge weight between two nodes
    pub fn get_edge_weight(&self, from_id: u64, to_id: u64) -> Option<EdgeWeight> {
        let from_index = self.get_node_index(from_id)?;
        let to_index = self.get_node_index(to_id)?;

        self.graph
            .find_edge(from_index, to_index)
            .and_then(|edge_index| self.graph.edge_weight(edge_index))
            .copied()
    }

    /// Checks if an edge exists between two nodes
    pub fn has_edge(&self, from_id: u64, to_id: u64) -> bool {
        self.get_edge_weight(from_id, to_id).is_some()
    }

    /// Gets all nodes at a specific layer
    pub fn get_nodes_at_layer(&self, layer: usize) -> Vec<u64> {
        self.graph
            .node_weights()
            .filter(|node| node.layer == layer)
            .map(|node| node.id)
            .collect()
    }

    /// Gets the maximum layer in the graph
    pub fn max_layer(&self) -> Option<usize> {
        self.graph
            .node_weights()
            .map(|node| node.layer)
            .max()
    }

    /// Adds a new node to the graph
    /// Returns true if the node was added successfully, false if a node with this ID already exists
    pub fn add_node(&mut self, node: VectorNode) -> bool {
        if self.contains_node(node.id) {
            return false;
        }

        let node_id = node.id;
        let node_index = self.graph.add_node(node);

        self.id_to_index.insert(node_id, node_index);
        self.index_to_id.insert(node_index, node_id);

        true
    }

    /// Adds a new node with the given parameters
    /// Returns true if the node was added successfully, false if a node with this ID already exists
    pub fn add_node_with_params(
        &mut self,
        id: u64,
        vector: Vec<f32>,
        layer: usize,
        max_connections: Option<usize>,
    ) -> bool {
        let max_connections = max_connections.unwrap_or(self.default_max_connections);
        let node = VectorNode::new(id, vector, layer, max_connections);
        self.add_node(node)
    }

    /// Removes a node from the graph along with all its edges
    /// Returns true if the node was removed, false if the node didn't exist
    pub fn remove_node(&mut self, id: u64) -> bool {
        if let Some(node_index) = self.get_node_index(id) {
            self.graph.remove_node(node_index);
            self.id_to_index.remove(&id);
            self.index_to_id.remove(&node_index);
            true
        } else {
            false
        }
    }

    /// Adds an edge between two nodes with the specified weight
    /// Returns true if the edge was added successfully, false if either node doesn't exist or edge already exists
    pub fn add_edge(&mut self, from_id: u64, to_id: u64, weight: EdgeWeight) -> bool {
        let from_index = match self.get_node_index(from_id) {
            Some(index) => index,
            None => return false,
        };

        let to_index = match self.get_node_index(to_id) {
            Some(index) => index,
            None => return false,
        };

        if self.graph.find_edge(from_index, to_index).is_some() {
            return false;
        }

        self.graph.add_edge(from_index, to_index, weight);
        true
    }

    /// Adds an edge between two nodes with calculated distance using the default metric
    /// Returns true if the edge was added successfully, false if either node doesn't exist or edge already exists
    pub fn add_edge_with_distance(&mut self, from_id: u64, to_id: u64, distance: f32) -> bool {
        let weight = EdgeWeight::new(distance, self.default_metric);
        self.add_edge(from_id, to_id, weight)
    }

    /// Removes an edge between two nodes
    /// Returns true if the edge was removed, false if the edge didn't exist
    pub fn remove_edge(&mut self, from_id: u64, to_id: u64) -> bool {
        let from_index = match self.get_node_index(from_id) {
            Some(index) => index,
            None => return false,
        };

        let to_index = match self.get_node_index(to_id) {
            Some(index) => index,
            None => return false,
        };

        if let Some(edge_index) = self.graph.find_edge(from_index, to_index) {
            self.graph.remove_edge(edge_index);
            true
        } else {
            false
        }
    }

    /// Updates the weight of an existing edge
    /// Returns true if the edge was updated, false if the edge doesn't exist
    pub fn update_edge_weight(&mut self, from_id: u64, to_id: u64, new_weight: EdgeWeight) -> bool {
        let from_index = match self.get_node_index(from_id) {
            Some(index) => index,
            None => return false,
        };

        let to_index = match self.get_node_index(to_id) {
            Some(index) => index,
            None => return false,
        };

        if let Some(edge_index) = self.graph.find_edge(from_index, to_index) {
            if let Some(edge_weight) = self.graph.edge_weight_mut(edge_index) {
                *edge_weight = new_weight;
                return true;
            }
        }
        false
    }

    /// Clears all nodes and edges from the graph
    pub fn clear(&mut self) {
        self.graph.clear();
        self.id_to_index.clear();
        self.index_to_id.clear();
    }
}
