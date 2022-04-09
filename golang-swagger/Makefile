NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP 			= l-sample
SERVER_BIN  	= ./bin/${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
IMAGE           = l-sample:canary
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)
RELEASE_VERSION = v0.0.1

default: help

## unittest: 运行单元测试
.PHONY: test
unittest:
	@CGO_ENABLED=0 go test -v $(shell go list ./pkg/...)

## integration-test: 运行集成测试
.PHONY: integration-test
integration-test:
	@CGO_ENABLED=0 go test -v $(shell go list ./test/...)

## swagger: 生成swagger文档
.PHONY: swagger
swagger:
	swag init --parseDependency -g ./pkg/swag/swag.go -o ./api/swagger

.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

