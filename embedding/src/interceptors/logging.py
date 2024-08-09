import logging
from typing import Any, Callable

from config import settings
from grpc import ServicerContext, StatusCode
from grpc_interceptor import ServerInterceptor
from grpc_interceptor.exceptions import GrpcException

logger = logging.getLogger(settings.embedding_server.name)


def status(code: StatusCode) -> str:
    """extract value string from grpc status code"""
    return code.value[1].upper()


class ExceptionLoggingInterceptor(ServerInterceptor):

    def intercept(
        self,
        method: Callable[..., Any],
        request_or_iterator: Any,
        context: ServicerContext,
        method_name: str,
    ) -> Any:
        """A logging interceptor that logs successful and failed exceptions
        and their reasons. Messages should reflect reports sent to client.

        :param method: The next interceptor, or method implementation.
        :param request_or_iterator: The RPC request, as a protobuf message.
        :param context: The ServicerContext pass by gRPC to the service.
        :param method_name: A string of the form "/protobuf.package.Service/Method"
        :return: the result of method(request_or_iterator, context)
        """
        method_name = method_name[method_name.rindex(".") + 1 :]
        try:
            response = method(request_or_iterator, context)
            logger.info(f"{method_name} {status(StatusCode.OK)}")
            return response
        except GrpcException as e:
            logger.error(f"{method_name} {status(e.status_code)} - Details: {e.details}")
            context.set_code(e.status_code)
            context.set_details(e.details)
            raise
