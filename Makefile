SHELL = /bin/bash

BUILD_IMAGE_SERVER  = golang:1.20.6
BUILD_IMAGE_WEB     = node:18


GOBIN = $(PWD)/.build/bin
GOPATH = $(PWD)/.build
GOBUILD = env CGO_ENABLED=1 GO111MODULE=on GOPROXY="https://goproxy.cn,direct" GOOS=linux GOARCH=amd64 go build
GOTEST = env GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go test

MODULE = $(shell sed 's/module //1p;d' go.mod)
MODULE_ROOT = $(shell dirname $(GOPATH)/src/$(MODULE))


.PHONY: prebuild apiserver hashts  clean test docker

all: apiserver 

prebuild:
	@mkdir -p $(MODULE_ROOT)
	@mkdir -p $(GOBIN)/txhash
	# @ln -sf $(PWD) $(MODULE_ROOT)

apiserver: prebuild hashts
	$(GOBUILD) -mod=mod -ldflags="$(GOLDFLAGS)" -o $(GOBIN)/$@ main.go
	@cp -rf config.docker.yaml $(GOBIN)/config.yaml
	@echo "server done"
	@echo "Run \"$(GOBIN)/$@\" to launch $@."
hashts:prebuild 
	@npm run --prefix ./txhash/ build
	@cp ./txhash/dist $(GOBIN)/txhash -r
	@echo "hashts done"

test:

clean:
	@for f in $(EXE_LIST); do rm -fv $(GOBIN)/$$f; done
docker:apiserver
	@echo "docker..."
	@manifest/docker/docker.sh apiserver
