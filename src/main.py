from config import settings
from proto import transformers_pb2_grpc
from service.server import Server
from service.transformers import SentenceTransformersService


def add_services(server: Server) -> None:
    """Add RPC services to server

    :param server: server instance to add services to
    """
    transformers_pb2_grpc.add_SentenceTransformersServicer_to_server(
        SentenceTransformersService(settings.transformers.model_list), server.instance
    )


if __name__ == "__main__":

    server = Server(
        address="0.0.0.0",
        port=settings.server.port,
        shutdown_period=settings.server.shutdown_period,
    )

    add_services(server)
    server.start()
