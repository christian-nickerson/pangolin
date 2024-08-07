import grpc
import pytest
from proto.transformers_pb2 import (  # type: ignore[attr-defined]
    InferenceRequest,
    InferenceResponse,
    ModelListRequest,
    ModelListResponse,
)
from proto.transformers_pb2_grpc import SentenceTransformersStub


def test_inference(server, address, model_name, lorem_embedding, lorem_ipsum) -> None:
    """test inference returns as expected"""
    with grpc.insecure_channel(address) as channel:
        stub = SentenceTransformersStub(channel)
        message = InferenceRequest(text=[lorem_ipsum], model_name=model_name)
        response: InferenceResponse = stub.Inference(message)
    embedding = [list(item.components) for item in response.embeddings]
    assert embedding == lorem_embedding


def test_inference_incorrect_model_name(server, address, lorem_ipsum) -> None:
    """test inference with wrong model name fails as expected"""
    with grpc.insecure_channel(address) as channel:
        stub = SentenceTransformersStub(channel)
        message = InferenceRequest(text=[lorem_ipsum], model_name="hello mum")
        with pytest.raises(grpc.RpcError) as e:
            _ = stub.Inference(message)
        assert e.value.code() == grpc.StatusCode.INVALID_ARGUMENT


def test_model_list(server, address, model_name) -> None:
    """test model_list returns as expected"""
    with grpc.insecure_channel(address) as channel:
        stub = SentenceTransformersStub(channel)
        message = ModelListRequest()
        response: ModelListResponse = stub.ModelList(message)
    model_list = [model for model in response.model_names]
    assert model_name in model_list
