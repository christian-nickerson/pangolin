from typing import List

from exceptions.server import ModelNotImplemented
from grpc import ServicerContext
from models.transformers import SentenceTransformerModels
from proto.embedding_pb2 import (  # type: ignore[attr-defined]
    InferenceRequest,
    InferenceResponse,
    ModelListRequest,
    ModelListResponse,
    Vector,
)
from proto.embedding_pb2_grpc import EmbeddingsServicer


class EmbeddingsService(EmbeddingsServicer):

    def __init__(self, model_list: List[str]):
        """Embeddings gRPC service

        :param model_list: List of models to make available
        """
        self.__transformers = SentenceTransformerModels(model_list)

    def Inference(self, request: InferenceRequest, context: ServicerContext) -> InferenceResponse:
        """Inference embeddings models

        :param request: inference request object
        :param context: generic context object
        :return: inference response object
        """
        if request.model_name not in self.__transformers.model_list:
            raise ModelNotImplemented(model_name=request.model_name)

        embeddings = self.__transformers.encode(list(request.text), request.model_name)
        message = [Vector(components=vector) for vector in embeddings]
        return InferenceResponse(embeddings=message)

    def ModelList(self, request: ModelListRequest, context: ServicerContext) -> ModelListResponse:
        """Return a list of available models

        :param request: model list request object
        :param context: generic context object
        :return: model list response object
        """
        return ModelListResponse(model_names=self.__transformers.model_list)
