SHELL=/bin/bash

.PHONY: proto

# 1. Compile server side code
# 2. Fix relative import error of generated code
proto:
	python -m grpc_tools.protoc -I=proto/ --python_out=embedding/src/proto --grpc_python_out=embedding/src/proto proto/*.proto
	cd embedding/src/proto && sed -i '' 's/^\(import.*pb2\)/from . \1/g' *.py
