from typing import List

import proto.transformers_pb2 as pb2
import proto.transformers_pb2_grpc as pb2_grpc
from models.transformers import SentenceTransformerModels


class SentenceTransformersService(pb2_grpc.SentenceTransformersServicer):
    def __init__(self, model_list: List[str]):
        """Sentence Transformers gRPC embedding service

        :param model_list: List of models to make available
        """
        self.transformers = SentenceTransformerModels(model_list)

    def Inference(self, request, context):
        """Inference sentence transformer model

        :param request: _description_
        :param context: _description_
        :return: _description_
        """
        embedding = self.transformers.encode(request.text, request.model_name)
        return pb2.InferenceResponse(embeddings=embedding)

    def ModelList(self, request, context):
        """Return a list of available models

        :param request: _description_
        :param context: _description_
        :return: _description_
        """
        return pb2.ModelListResponse(model_names=self.transformers.model_list)
