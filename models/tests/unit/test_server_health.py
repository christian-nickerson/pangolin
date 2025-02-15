import grpc
from grpc_health.v1 import health_pb2, health_pb2_grpc


def test_health_check(server, address) -> None:
    """test inference returns as expected"""
    with grpc.insecure_channel(address) as channel:
        stub = health_pb2_grpc.HealthStub(channel)
        message = health_pb2.HealthCheckRequest(service="")
        response: health_pb2.HealthCheckResponse = stub.Check(message)
    assert response.status == health_pb2.HealthCheckResponse.SERVING
