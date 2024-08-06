from typing import List

import pytest
from exceptions.embedding import ModelRemoteImportError
from models.transformers import SentenceTransformerModels
from sentence_transformers import SentenceTransformer


@pytest.fixture
def model_name() -> str:
    return "all-MiniLM-L6-v2"


@pytest.fixture
def lorem_ipsum() -> str:
    """lorem ipsum test text"""
    with open("tests/fixtures/lorem.txt", "r") as file:
        return file.read()


@pytest.fixture
def lorem_embedding(model_name, lorem_ipsum) -> List[float]:
    """return lorem ipsum embedding"""
    st = SentenceTransformer(model_name)
    return st.encode(lorem_ipsum, convert_to_numpy=True).tolist()


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
    embedding = stm.encode(lorem_ipsum, model_name)
    assert embedding == lorem_embedding
