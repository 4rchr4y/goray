RAY_DIR := .ray

PROTO_GO_PKG := "./internal/proto/"
PROTO_DIR := "protocol"
PROTO_DRIVER_VERSION := 1.0
PROTO_COMPONENT_VERSION := 1.0


.PHONY: protoc
protoc:
	@export PATH="${PATH}:$(shell go env GOPATH)/bin" ; \
	protoc -I ./${PROTO_DIR}/ \
		--go_out=${PROTO_GO_PKG} --go_opt=paths=source_relative \
		--go-grpc_out=${PROTO_GO_PKG} --go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/protopkg/schema.proto \
		${PROTO_DIR}/protopkg/diagnostic.proto \
		${PROTO_DIR}/protodriver/driver${PROTO_DRIVER_VERSION}.proto \
		${PROTO_DIR}/protocomponent/component${PROTO_COMPONENT_VERSION}.proto


.PHONY: count
count:
	@ find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.PHONY: noop_driver
noop_driver:
	go build -o ./.ray/driver/github.com/4rchr4y/ray-noop-driver@v0.0.1 ./example/noop-driver/main


.PHONY: dummy_component
dummy_component:
	go build -o ./.ray/component/github.com/4rchr4y/ray-dummy-component@v0.0.1 ./example/dummy-component/main


