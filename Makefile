CFLAGS=-g
export CFLAGS

.PHONY: build run build-all release
build:
	go build -o bin/main main.go

.PHONY: run the program
run:
	go run main.go

.PHONY: build-all build for all platforms
build-all:
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

.PHONY: release releases a new version
release:
	git tag -a v$(v) -m "Release v$(v)"
	git push origin v$(v)
	git push