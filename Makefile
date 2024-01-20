#Dockerfile vars

#vars
IMAGENAME=go-proxy
TAG=latest
BUILDDATE=${shell date -u +%Y-%m-%dT%H:%M:%SZ}
IMAGEFULLNAME=avhost/${IMAGENAME}
BRANCH=${shell git symbolic-ref --short HEAD}

.PHONY: help build bootstrap all docs publish push version

help:
	    @echo "Makefile arguments:"
	    @echo ""
	    @echo "Makefile commands:"
			@echo "push"
	    @echo "build"
			@echo "build-bin"
	    @echo "all"
			@echo "docs"
			@echo "publish"
			@echo "version"
			@echo ${TAG}

.DEFAULT_GOAL := all

build:
	@echo ">>>> Build docker image"
	@docker build --build-arg TAG=${TAG} --build-arg BUILDDATE=${BUILDDATE} -t ${IMAGEFULLNAME}:latest .


build-bin:
	@echo ">>>> Build binary"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.BuildVersion=${BUILDDATE} -X main.GitVersion=${TAG} -X main.VersionURL=${VERSION_URL} -extldflags \"-static\"" .

push:
	@echo ">>>> Build docker image"
	@docker buildx build --push --build-arg TAG=${TAG} --build-arg BUILDDATE=${BUILDDATE} -t ${IMAGEFULLNAME}:latest .
	@docker buildx build --push --build-arg TAG=${TAG} --build-arg BUILDDATE=${BUILDDATE} -t ${IMAGEFULLNAME}:${TAG} .


update-gomod:
	go get -u

seccheck:
	gosec --exclude G104 --exclude-dir ./vendor ./... 

sboom:
	syft dir:. > sbom.txt
	syft dir:. -o json > sbom.json

all: build seccheck sboom push
