import proto.embedding_pb2_grpc as embeddings
from config import settings
from logger import create_logger
from service.models import EmbeddingsService
from service.server import Server


def add_services(server: Server) -> None:
    """Add RPC services to server

    :param server: server instance to add services to
    """
    embeddings.add_EmbeddingsServicer_to_server(EmbeddingsService(settings.transformers.model_list), server.instance)


if __name__ == "__main__":

    loggger = create_logger(settings.server.embeddings.name)

    server = Server(
        address="[::]",
        port=settings.server.embeddings.port,
        shutdown_period=settings.server.embeddings.shutdown_period,
        max_worker_threads=settings.server.embeddings.worker_threads,
    )

    add_services(server)
    server.start()
