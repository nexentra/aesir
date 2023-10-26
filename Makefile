CFLAGS=-g
SHELL=/bin/bash
GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

BIN=dist/bin/aesir
SRC=./cmd/aesir
SRC_API=./cmd/aesir-api

LDFLAGS=-w -s -X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}
export CFLAGS
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64

.PHONY: all clean vet build docker-image run release
vet:
	go vet ./...

build: vet
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${SRC}"

docker-build:
	docker build \
		--no-cache=true \
		--build-arg BUILD_REV="${GIT_HASH}" \
		--build-arg BUILD_DATE="${BUILD_DATE}" \
		-t nexentra/aesir:latest \
		-f Dockerfile.scratch .

docker-run:
	if [ -d $(path) ]; then \
		docker run -it --rm -i nexentra/aesir; \
	else \
		docker run -it --rm -v $(path):$(path) nexentra/aesir $(path); \
	fi

docker-push: clean docker-build
	docker push nexentra/aesir:${v}

run:
	if [ -d $(path) ]; then \
		go run ${SRC}; \
	else \
		go run ${SRC} $(path); \
	fi

run-api:
	if [ -d $(port) ]; then \
		go run ${SRC_API}; \
	else \
		go run ${SRC_API} -port=$(port); \
	fi

gen:
	go generate ./...

vis-graph:
	go run -mod=mod ariga.io/entviz ./ent/schema 

clean:
	rm -rf dist

release:
	git tag -a v$(v) -m "Release v$(v)"
	git push origin v$(v)
	git push