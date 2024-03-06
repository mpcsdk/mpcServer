SHELL = /bin/bash
USER_ID ?= $(shell id -u)
GROUP_ID ?= $(shell id -g)
BUILD_IMAGE_SERVER  = golang:1.20
BUILD_IMAGE_NODE = node:18
PROJECT_NAME        = "mpcServer"
Image_NAME = "mpcserver"
ifeq ($(TAGS_OPT),)
TAGS_OPT            = latest
else
endif

build:  build-server
	docker run -u $(USER_ID):$(GROUP_ID) --name build-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_NODE} make build-local

build-server:
	docker run  -e USER_ID=$(USER_ID)  -e GROUP_ID=$(GROUP_ID)  --name build-server-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-server-local

build-local:
	if [ -d "build" ];then rm -rf build; else echo "build OK!"; fi \
	&& mkdir build \
	&& if [ -f "/.dockerenv" ];then echo "dockerenv OK!"; else  make build-server-local; fi \
	&& cp ./mpcServer build/ && cp -r ./config.docker.yaml build/config.yaml \
	&& make build-hash


build-server-local:
	if [ -f "mpcServer" ];then rm -rf mpcServer; else echo "mpcServer OK!"; fi 
	 go env -w GOPROXY=https://goproxy.cn,direct 
	 go env -w CGO_ENABLED=1 && go env  && go mod tidy 
	 git config --global --add safe.directory /go/src/mpcServer
	 go build -ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${TAGS_OPT}" -v  -o mpcServer
	 chown -R $(USER_ID):$(GROUP_ID) ./mpcServer

build-hash:
	rm ./build/utility -rf && mkdir -p ./build/utility/txhash/dist \
	&& npm run --prefix ./utility/txhash/ build \
	&& cp ./utility/txhash/dist/* ./build/utility/txhash/dist/ -rf \
	&& cp ./utility/txhash/protobuf ./build/utility/txhash/protobuf -rf \
	&& echo "hashts done"

image: build
	docker build -t ${Image_NAME}:${TAGS_OPT} -f manifest/docker/Dockerfile .

