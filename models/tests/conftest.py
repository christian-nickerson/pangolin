from typing import Iterator, List

import pytest
from config import settings
from logger import create_logger
from main import add_services
from sentence_transformers import SentenceTransformer
from service.server import Server


@pytest.fixture(scope="session")
def port() -> int:
    """grpc server port"""
    return 50051


@pytest.fixture(scope="session")
def address(port) -> str:
    """grpc server port"""
    return f"localhost:{port}"


@pytest.fixture(scope="session")
def server(port) -> Iterator[Server]:
    """start grpc server, returns address"""
    _ = create_logger(settings.server.embeddings.name)
    server = Server(port=port, shutdown_period=5)
    add_services(server)
    server.start(False)
    yield server
    server.stop()


@pytest.fixture
def model_name() -> str:
    return "all-MiniLM-L6-v2"


@pytest.fixture
def lorem_ipsum() -> str:
    """lorem ipsum test text"""
    with open("embedding/tests/fixtures/lorem.txt", "r") as file:
        return file.read()


@pytest.fixture
def lorem_embedding(model_name, lorem_ipsum) -> List[float]:
    """return lorem ipsum embedding"""
    st = SentenceTransformer(model_name)
    return st.encode([lorem_ipsum], convert_to_numpy=True).tolist()
