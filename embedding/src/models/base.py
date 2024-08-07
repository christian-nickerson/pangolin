from abc import ABC, abstractmethod
from typing import List


class EmbeddingModels(ABC):

    @abstractmethod
    def encode(self, text: List[str], model_name: str) -> List[List[float]]:
        raise NotImplementedError()
