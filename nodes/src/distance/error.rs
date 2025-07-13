use std::fmt;

#[allow(dead_code)]
#[derive(Debug, PartialEq)]
pub enum DistanceError {
    DimensionMismatch { len1: usize, len2: usize },
    ZeroMagnitude,
}

impl fmt::Display for DistanceError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            DistanceError::DimensionMismatch { len1, len2 } => {
                write!(f, "Vector dimension mismatch: {} != {}", len1, len2)
            }
            DistanceError::ZeroMagnitude => {
                write!(
                    f,
                    "Cannot compute cosine similarity with zero magnitude vector"
                )
            }
        }
    }
}

impl std::error::Error for DistanceError {}
