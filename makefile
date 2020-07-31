PROJECT_NAME=robusta
BUILD_VERSION=$(shell cat VERSION)
DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
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
	cd deployment/docker && BUILD_VERSION=$(BUILD_VERSION) docker-compose up

docker_prebuild: vet web_build build
	mkdir -p deployment/docker/web/dist
	mkdir -p deployment/docker/configs
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin deployment/docker/$(PROJECT_NAME).bin; \
	cp -R web/dist deployment/docker/web/; \
	cp -R configs deployment/docker/;

docker_build:
	cd deployment/docker; \
	docker build -t $(DOCKER_IMAGE) .;

docker_postbuild:
	cd deployment/docker; \
	rm -rf $(PROJECT_NAME).bin 2> /dev/null;\
	rm -rf web 2> /dev/null; \
	rm -rf configs 2> /dev/null;