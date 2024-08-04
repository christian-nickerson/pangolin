from typing import List

from grpc import ServicerContext
from grpc_interceptor.exceptions import Internal

import proto.transformers_pb2 as pb2
import proto.transformers_pb2_grpc as pb2_grpc
from models.transformers import SentenceTransformerModels


class SentenceTransformersService(pb2_grpc.SentenceTransformersServicer):

    def __init__(self, model_list: List[str]):
        """Sentence Transformers gRPC embedding service

        :param model_list: List of models to make available
        """
        self.__transformers = SentenceTransformerModels(model_list)

    def Inference(self, request: pb2.InferenceRequest, context: ServicerContext):
        """Inference sentence transformer model

        :param request: inference request object
        :param context: generic context object
        :return: inference response object
        """
        embedding = self.__transformers.encode(request.text, request.model_name)
        return pb2.InferenceResponse(embeddings=embedding)

    def ModelList(self, request: pb2.ModelListRequest, context: ServicerContext):
        """Return a list of available models

        :param request: model list request object
        :param context: generic context object
        :return: model list respone object
        """
        return pb2.ModelListResponse(model_names=self.__transformers.model_list)
