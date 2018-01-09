APP?=hook
PORT?=8080
PROJECT?=github.com/gotoolkit/hook
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}

build: clean
	go build -ldflags "-s -w \
		-X ${PROJECT}/pkg/version.Release=${RELEASE} \
		-X ${PROJECT}/pkg/version.Commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.BuildTime=${BUILD_TIME} " \
		-o ${APP}

run: build
	PORT=${PORT} ./${APP}

test:
	go test -v -race ./...