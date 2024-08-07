import pytest
from exceptions.embedding import ModelRemoteImportError
from models.transformers import SentenceTransformerModels


def test_model_import_and_list(model_name) -> None:
    """test models can import and return the model_list property"""
    stm = SentenceTransformerModels([model_name])
    assert stm.model_list == [model_name]


def test_model_import_failure() -> None:
    """test model import fails as expected"""
    model_names = ["hi mum"]
    with pytest.raises(ModelRemoteImportError):
        SentenceTransformerModels(model_names)


def test_encode(model_name, lorem_ipsum, lorem_embedding) -> None:
    """test embedding encoding is as expected"""
    stm = SentenceTransformerModels([model_name])
    embedding = stm.encode([lorem_ipsum], model_name)
    assert embedding == lorem_embedding
