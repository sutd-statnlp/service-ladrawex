GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

all: test

run: 
	bash script/run.sh

run-prod: 
	env SL_MODE=prod make run

test:
	bash script/test.sh

test-cov:
	bash script/test-coverage.sh	

build:
	bash script/build.sh

build-multi:
	bash script/build-multi.sh

docker:
	bash script/docker-build.sh

docker-push:
	bash script/docker-push.sh