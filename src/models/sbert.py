from typing import Dict, List

from sentence_transformers import SentenceTransformer
from sentence_transformers.util import is_sentence_transformer_model

from errors.embedding import ModelRemoteImportError
from models.base import EmbeddingModels


class SentenceTransformerModels(EmbeddingModels):
    def __init__(self, model_list: List[str]) -> None:
        """_summary_"""
        self.registry = self._load_models(model_list)

    def encode(self, text: List[str], model_name: str) -> List[float]:
        """Inference model with text to generate an embedding

        :param text: text used to generate embedding
        :param model_name: model name to use to generate embedding
        :return: embedding
        """
        embedding = self.registry[model_name].encode(text)
        return embedding.tolist()

    @staticmethod
    def _load_models(model_list: List[str]) -> Dict[str, SentenceTransformer]:
        """Load all models into a dictionary

        :param model_list: list of all model names to import
        :return: model registry
        """
        registry = {}
        for model_name in model_list:
            if is_sentence_transformer_model(model_name):
                registry[model_name] = SentenceTransformer(model_name)
            else:
                raise ModelRemoteImportError(model_name, "huggingface.co/sentence-transformers")
        return registry

    @property
    def model_list(self) -> List[str]:
        """List all models that have been loaded

        :return: List of available models
        """
        return list(self.registry.keys())
