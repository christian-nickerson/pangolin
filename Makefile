SHELL=/bin/bash

.PHONY: proto

proto:
	python -m grpc_tools.protoc -I=proto/ --python_out=src/proto --grpc_python_out=src/proto proto/*.proto
