SHELL=/bin/bash

.PHONY: proto

proto:
	python -m grpc_tools.protoc -I=proto/ --python_out=embedding/src/proto --grpc_python_out=embedding/src/proto proto/*.proto
