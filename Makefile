SHELL=/bin/bash

.PHONY: proto

# 1. Compile go client side code
# 2. Compile python server side code (seperate due to python protoc bug)
# 2. Fix relative import error of generated code in python
proto:

	protoc -I=proto \
		--go_out=api/internal/proto \
		--go_opt=paths=source_relative  \
		--go-grpc_out=api/internal/proto  \
		--go-grpc_opt=paths=source_relative  \
		proto/*.proto

	python -m grpc_tools.protoc -I=proto \
		--python_out=embedding/src/proto \
		--pyi_out=embedding/src/proto \
		--grpc_python_out=embedding/src/proto \
		proto/*.proto

	cd embedding/src/proto && sed -i 's/^\(import.*pb2\)/from . \1/g' *.py
