PROJECT_NAME=gotripe
BUILD_VERSION=$(shell cat VERSION)
DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
WEB_DOCKER_IMAGE=$(PROJECT_NAME):client
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on


all: mod_tidy fmt vet 

build:
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin .

vet:
	$(GO_BUILD_ENV) go vet $(GO_FILES)

fmt:
	$(GO_BUILD_ENV) go fmt $(GO_FILES)

mod_tidy:
	$(GO_BUILD_ENV) go mod tidy

compose_prod: docker
	BUILD_VERSION=$(BUILD_VERSION) docker-compose up

docker_build:
	docker build -t $(DOCKER_IMAGE) .;

docker_buildweb:
	cd web; \
	docker build -t $(WEB_DOCKER_IMAGE) .; \
	cd ../;

docker: vet build docker_build 