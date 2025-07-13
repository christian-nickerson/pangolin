#[allow(dead_code)]
pub fn l2_distance(
    vec1: &[f64],
    vec2: &[f64],
) -> Result<f64, crate::distance::error::DistanceError> {
    if vec1.len() != vec2.len() {
        return Err(crate::distance::error::DistanceError::DimensionMismatch {
            len1: vec1.len(),
            len2: vec2.len(),
        });
    }

    let squared_distance: f64 = vec1
        .iter()
        .zip(vec2.iter())
        .map(|(a, b)| (a - b).powi(2))
        .sum();

    Ok(squared_distance.sqrt())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_l2_distance() {
        let vec1 = vec![2.0, 3.0];
        let vec2 = vec![5.0, 7.0];

        let result = l2_distance(&vec1, &vec2);
        assert!(result.is_ok());

        let distance = result.unwrap();
        assert_eq!(distance, 5.0);
    }

    #[test]
    fn test_l2_distance_dimension_mismatch() {
        let vec1 = vec![1.0, 2.0, 3.0];
        let vec2 = vec![4.0, 5.0];

        let result = l2_distance(&vec1, &vec2);
        assert!(result.is_err());

        match result.unwrap_err() {
            crate::distance::error::DistanceError::DimensionMismatch { len1, len2 } => {
                assert_eq!(len1, vec1.len());
                assert_eq!(len2, vec2.len());
            }
            _ => panic!("Expected DimensionMismatch error"),
        }
    }
}
