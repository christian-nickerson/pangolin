import proto.transformers_pb2_grpc as transformers
from config import settings
from logger import create_logger
from service.server import Server
from service.transformers import SentenceTransformersService


def add_services(server: Server) -> None:
    """Add RPC services to server

    :param server: server instance to add services to
    """
    transformers.add_SentenceTransformersServicer_to_server(
        SentenceTransformersService(settings.transformers.model_list), server.instance
    )


if __name__ == "__main__":

    loggger = create_logger(settings.server.embedding.name)

    server = Server(
        address="[::]",
        port=settings.server.embedding.port,
        shutdown_period=settings.server.embedding.shutdown_period,
        max_worker_threads=settings.server.embedding.worker_threads,
    )

    add_services(server)
    server.start()
