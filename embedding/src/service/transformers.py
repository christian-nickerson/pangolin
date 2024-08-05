from typing import List

from grpc import ServicerContext
from models.transformers import SentenceTransformerModels
from proto.transformers_pb2 import (  # type: ignore[attr-defined]
    InferenceRequest,
    InferenceResponse,
    ModelListRequest,
    ModelListResponse,
)
from proto.transformers_pb2_grpc import SentenceTransformersServicer


class SentenceTransformersService(SentenceTransformersServicer):

    def __init__(self, model_list: List[str]):
        """Sentence Transformers gRPC embedding service

        :param model_list: List of models to make available
        """
        self.__transformers = SentenceTransformerModels(model_list)

    def Inference(self, request: InferenceRequest, context: ServicerContext) -> InferenceResponse:
        """Inference sentence transformer model

        :param request: inference request object
        :param context: generic context object
        :return: inference response object
        """
        embedding = self.__transformers.encode(request.text, request.model_name)
        return InferenceResponse(embeddings=embedding)

    def ModelList(self, request: ModelListRequest, context: ServicerContext) -> ModelListResponse:
        """Return a list of available models

        :param request: model list request object
        :param context: generic context object
        :return: model list respone object
        """
        return ModelListResponse(model_names=self.__transformers.model_list)
