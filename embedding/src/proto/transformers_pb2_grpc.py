# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import transformers_pb2 as transformers__pb2

GRPC_GENERATED_VERSION = "1.65.4"
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = "1.66.0"
SCHEDULED_RELEASE_DATE = "August 6, 2024"
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower

    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f"The grpc package installed is at version {GRPC_VERSION},"
        + f" but the generated code in transformers_pb2_grpc.py depends on"
        + f" grpcio>={GRPC_GENERATED_VERSION}."
        + f" Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}"
        + f" or downgrade your generated code using grpcio-tools<={GRPC_VERSION}."
        + f" This warning will become an error in {EXPECTED_ERROR_RELEASE},"
        + f" scheduled for release on {SCHEDULED_RELEASE_DATE}.",
        RuntimeWarning,
    )


class SentenceTransformersStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Inference = channel.unary_unary(
            "/EmbeddingService.SentenceTransformers/Inference",
            request_serializer=transformers__pb2.InferenceRequest.SerializeToString,
            response_deserializer=transformers__pb2.InferenceResponse.FromString,
            _registered_method=True,
        )
        self.ModelList = channel.unary_unary(
            "/EmbeddingService.SentenceTransformers/ModelList",
            request_serializer=transformers__pb2.ModelListRequest.SerializeToString,
            response_deserializer=transformers__pb2.ModelListResponse.FromString,
            _registered_method=True,
        )


class SentenceTransformersServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Inference(self, request, context):
        """Inference an embedding model"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")

    def ModelList(self, request, context):
        """Model list"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_SentenceTransformersServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "Inference": grpc.unary_unary_rpc_method_handler(
            servicer.Inference,
            request_deserializer=transformers__pb2.InferenceRequest.FromString,
            response_serializer=transformers__pb2.InferenceResponse.SerializeToString,
        ),
        "ModelList": grpc.unary_unary_rpc_method_handler(
            servicer.ModelList,
            request_deserializer=transformers__pb2.ModelListRequest.FromString,
            response_serializer=transformers__pb2.ModelListResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "EmbeddingService.SentenceTransformers", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers("EmbeddingService.SentenceTransformers", rpc_method_handlers)


# This class is part of an EXPERIMENTAL API.
class SentenceTransformers(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Inference(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/EmbeddingService.SentenceTransformers/Inference",
            transformers__pb2.InferenceRequest.SerializeToString,
            transformers__pb2.InferenceResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True,
        )

    @staticmethod
    def ModelList(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/EmbeddingService.SentenceTransformers/ModelList",
            transformers__pb2.ModelListRequest.SerializeToString,
            transformers__pb2.ModelListResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True,
        )
