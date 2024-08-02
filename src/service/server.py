import signal
from concurrent import futures

import grpc


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
        :param shutdown_period: Seconds to wait for RPC processes to finish before shutdown
        """
        self.__address = address
        self.__port = port
        self.__shutdown_period = shutdown_period
        self.__shutdown_config()
        self.__server = grpc.server(futures.ThreadPoolExecutor(max_workers=max_worker_threads))

    def start(self) -> None:
        """Starts the grpc server"""
        endpoint = f"{self.__address}:{self.__port}"
        self.__server.add_insecure_port(endpoint)
        self.__server.start()
        self.__server.wait_for_termination()

    def __shutdown_config(self) -> None:
        """Handle signal interrupts to gracefully shutdown server"""
        signal.signal(signal.SIGINT, self.stop)
        signal.signal(signal.SIGTERM, self.stop)
        signal.signal(signal.SIGKILL, self.stop)

    def stop(self, *args) -> None:
        """Stops the server gracefully"""
        self.__server.stop(self.__shutdown_period)

    @property
    def instance(self):
        """GRPC server instance"""
        return self.__server
