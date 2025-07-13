#[cfg(test)]
mod integration_tests {
    use crate::hnsw::build::{HnswIndex, HnswConfig};
    use crate::hnsw::graph::DistanceMetric;

    #[test]
    fn test_hnsw_euclidean_integration() {
        let config = HnswConfig {
            m: 8,
            max_m: 8,
            max_m_l: 16,
            ml: 1.0 / (2.0_f64).ln(),
            ef_construction: 100,
            metric: DistanceMetric::Euclidean,
        };

        let mut index = HnswIndex::new(config);

        // Create test vectors in 2D space
        let vectors = vec![
            (1, vec![0.0, 0.0]),    // Origin
            (2, vec![1.0, 0.0]),    // Right
            (3, vec![0.0, 1.0]),    // Up
            (4, vec![1.0, 1.0]),    // Diagonal
            (5, vec![-1.0, 0.0]),   // Left
            (6, vec![0.0, -1.0]),   // Down
            (7, vec![2.0, 0.0]),    // Far right
            (8, vec![0.0, 2.0]),    // Far up
        ];

        // Build index
        assert!(index.build_from_vectors(vectors).is_ok());
        assert_eq!(index.len(), 8);

        // Test search near origin - should find origin first
        let results = index.search(&[0.1, 0.1], 3, 50).unwrap();
        assert_eq!(results.len(), 3);
        assert_eq!(results[0].0, 1); // Origin should be closest

        // Test search near (1, 0) - should find vector 2 first
        let results = index.search(&[0.9, 0.1], 2, 50).unwrap();
        assert_eq!(results.len(), 2);
        assert_eq!(results[0].0, 2); // (1, 0) should be closest

        // Test search for exact match
        let results = index.search(&[1.0, 1.0], 1, 50).unwrap();
        assert_eq!(results.len(), 1);
        assert_eq!(results[0].0, 4); // (1, 1) should be exact match
        assert!(results[0].1 < 0.001); // Distance should be very small
    }

    #[test]
    fn test_hnsw_cosine_integration() {
        let config = HnswConfig {
            m: 8,
            max_m: 8,
            max_m_l: 16,
            ml: 1.0 / (2.0_f64).ln(),
            ef_construction: 100,
            metric: DistanceMetric::Cosine,
        };

        let mut index = HnswIndex::new(config);

        // Create test vectors with different magnitudes but similar directions
        let vectors = vec![
            (1, vec![1.0, 0.0, 0.0]),      // Unit vector along X
            (2, vec![2.0, 0.0, 0.0]),      // 2x along X (same direction)
            (3, vec![0.0, 1.0, 0.0]),      // Unit vector along Y
            (4, vec![0.0, 0.0, 1.0]),      // Unit vector along Z
            (5, vec![1.0, 1.0, 0.0]),      // 45 degrees in XY plane
            (6, vec![1.0, 0.0, 1.0]),      // 45 degrees in XZ plane
        ];

        // Build index
        assert!(index.build_from_vectors(vectors).is_ok());
        assert_eq!(index.len(), 6);

        // Test search for vector similar to (1, 0, 0)
        // Should find vectors 1 and 2 as most similar (same direction)
        let results = index.search(&[3.0, 0.0, 0.0], 2, 50).unwrap();
        assert_eq!(results.len(), 2);
        // Both vector 1 and 2 should have very low cosine distance (high similarity)
        assert!(results[0].1 < 0.1);
        assert!(results[1].1 < 0.1);
    }

    #[test]
    fn test_hnsw_empty_index() {
        let index = HnswIndex::with_default_config();

        assert!(index.is_empty());
        assert_eq!(index.len(), 0);
        assert!(index.entry_point().is_none());

        // Search on empty index should return empty results
        let results = index.search(&[1.0, 2.0, 3.0], 5, 10).unwrap();
        assert!(results.is_empty());
    }

    #[test]
    fn test_hnsw_single_vector() {
        let mut index = HnswIndex::with_default_config();

        let vector = vec![1.0, 2.0, 3.0];
        assert!(index.insert(42, vector.clone()).is_ok());

        assert!(!index.is_empty());
        assert_eq!(index.len(), 1);
        assert_eq!(index.entry_point(), Some(42));

        // Search should return the single vector
        let results = index.search(&[1.1, 2.1, 3.1], 1, 10).unwrap();
        assert_eq!(results.len(), 1);
        assert_eq!(results[0].0, 42);
    }

    #[test]
    fn test_hnsw_duplicate_insertion() {
        let mut index = HnswIndex::with_default_config();

        let vector = vec![1.0, 2.0, 3.0];

        // First insertion should succeed
        assert!(index.insert(1, vector.clone()).is_ok());
        assert_eq!(index.len(), 1);

        // Second insertion with same ID should not increase size
        assert!(index.insert(1, vector.clone()).is_ok());
        assert_eq!(index.len(), 1);
    }
}
