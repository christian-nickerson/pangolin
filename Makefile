SHELL=/bin/bash

.PHONY: proto

proto:

	protoc -I=proto \
		--go_out=control/internal/proto \
		--go_opt=paths=source_relative  \
		--go-grpc_out=control/internal/proto  \
		--go-grpc_opt=paths=source_relative  \
		proto/*.proto

	python -m grpc_tools.protoc -I=proto \
		--python_out=models/src/proto \
		--pyi_out=models/src/proto \
		--grpc_python_out=models/src/proto \
		proto/*.proto

	cd models/src/proto && sed -i 's/^\(import.*pb2\)/from . \1/g' *.py

docker-build:

	docker build -f Dockerfile.control -t pangolin-control:latest .
	docker build -f Dockerfile.mdoels -t pangolin-mdoels:latest .
