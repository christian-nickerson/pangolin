#[allow(dead_code)]
fn dot_product(vec1: &[f64], vec2: &[f64]) -> f64 {
    vec1.iter().zip(vec2.iter()).map(|(a, b)| a * b).sum()
}

#[allow(dead_code)]
fn magnitude(vec: &[f64]) -> f64 {
    vec.iter().map(|x| x * x).sum::<f64>().sqrt()
}

#[allow(dead_code)]
pub fn cosine_similarity(
    vec1: &[f64],
    vec2: &[f64],
) -> Result<f64, crate::distance::error::DistanceError> {
    if vec1.len() != vec2.len() {
        return Err(crate::distance::error::DistanceError::DimensionMismatch {
            len1: vec1.len(),
            len2: vec2.len(),
        });
    }

    let dot = dot_product(vec1, vec2);
    let mag1 = magnitude(vec1);
    let mag2 = magnitude(vec2);

    if mag1 == 0.0 || mag2 == 0.0 {
        return Err(crate::distance::error::DistanceError::ZeroMagnitude);
    }

    Ok(dot / (mag1 * mag2))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cosine_similarity() {
        let vec1 = vec![1.0, 2.0, 3.0];
        let vec2 = vec![4.0, 5.0, 6.0];

        let result = cosine_similarity(&vec1, &vec2);
        assert!(result.is_ok());

        let similarity = result.unwrap();
        let expected = 32.0 / (14.0_f64.sqrt() * 77.0_f64.sqrt());
        assert!((similarity - expected).abs() < 1e-10);
    }

    #[test]
    fn test_cosine_similarity_dimension_mismatch() {
        let vec1 = vec![1.0, 2.0, 3.0];
        let vec2 = vec![4.0, 5.0];

        let result = cosine_similarity(&vec1, &vec2);
        assert!(result.is_err());

        match result.unwrap_err() {
            crate::distance::error::DistanceError::DimensionMismatch { len1, len2 } => {
                assert_eq!(len1, 3);
                assert_eq!(len2, 2);
            }
            _ => panic!("Expected DimensionMismatch error"),
        }
    }

    #[test]
    fn test_cosine_similarity_zero_magnitude() {
        let vec1 = vec![0.0, 0.0, 0.0];
        let vec2 = vec![1.0, 2.0, 3.0];

        let result = cosine_similarity(&vec1, &vec2);
        assert!(result.is_err());

        match result.unwrap_err() {
            crate::distance::error::DistanceError::ZeroMagnitude => {}
            _ => panic!("Expected ZeroMagnitude error"),
        }
    }
}
