import logging
import signal
from concurrent import futures

import grpc
from config import settings
from grpc_health.v1 import health, health_pb2, health_pb2_grpc
from interceptors.logging import ExceptionLoggingInterceptor

logger = logging.getLogger(settings.embedding_server.name)


class Server:

    def __init__(
        self,
        address="[::]",
        port=50051,
        max_worker_threads: int = 10,
        shutdown_period: int = 5,
    ):
        """GRPC Server for running a gRPC API server

        :param address: Address to handle requests on, defaults to "[::]"
        :param port: Port to handle requests from, defaults to 50051
        :param max_worker_threads: Number of threads to assign to server
        :param shutdown_period: Seconds to wait for RPC processes to finish before shutdown
        """
        self.__address = address
        self.__port = port
        self.__shutdown_period = shutdown_period
        self.__server = grpc.server(
            thread_pool=futures.ThreadPoolExecutor(max_workers=max_worker_threads),
            interceptors=[ExceptionLoggingInterceptor()],
        )

        self.__health_check_config(max_worker_threads)
        self.__shutdown_config()

    def start(self, wait_for_termination: bool = True) -> None:
        """Starts the grpc server

        :param wait_for_termination: enter server into wait for termination. Primarily used for test purposes.
        """
        endpoint = f"{self.__address}:{self.__port}"
        self.__server.add_insecure_port(endpoint)
        self.__health.set("", health_pb2.HealthCheckResponse.SERVING)
        self.__server.start()
        logger.info(f"serving on {self.__address} port {self.__port}")
        if wait_for_termination:
            self.__server.wait_for_termination()

    def __shutdown_config(self) -> None:
        """Handle signal interrupts to gracefully shutdown server"""
        signal.signal(signal.SIGINT, self.stop)
        signal.signal(signal.SIGTERM, self.stop)

    def __health_check_config(self, max_worker_threads: int) -> None:
        """Configure health check endpoints for server

        :param max_worker_threads: Number of threads to assign to server
        """
        self.__health = health.HealthServicer(
            experimental_non_blocking=True,
            experimental_thread_pool=futures.ThreadPoolExecutor(max_workers=max_worker_threads),
        )
        health_pb2_grpc.add_HealthServicer_to_server(self.__health, self.__server)

    def stop(self, *args) -> None:
        """Stops the server gracefully"""
        logger.debug("server stopping...")
        self.__health.enter_graceful_shutdown()
        self.__server.stop(self.__shutdown_period)
        logger.info("server shutdown safely")

    @property
    def instance(self):
        """GRPC server instance"""
        return self.__server
