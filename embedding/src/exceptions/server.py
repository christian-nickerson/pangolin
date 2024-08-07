from grpc import StatusCode
from grpc_interceptor.exceptions import GrpcException


class ModelNotImplemented(GrpcException):

    status_code: StatusCode = StatusCode.INVALID_ARGUMENT
    details: str = "{name} is not implemented. Please see implemented models from ModelList method."

    def __init__(
        self,
        model_name: str,
    ) -> None:
        """Exception raised when inference models is not implemented on embedding server.

        :param model_name: name of model not implemented
        """
        self.details = self.details.format(name=model_name)
