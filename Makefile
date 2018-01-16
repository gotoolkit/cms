APP?=cms
PORT?=8000
MYSQL_DATABASE?=root:root@tcp(docker.for.mac.localhost:3306)/sme?charset=utf8mb4&parseTime=true
PROJECT?=github.com/gotoolkit/cms
TELEGRAM_HORN_URL?=https://integram.org/crVSWQfVmah

RELEASE?=0.1.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64

CONTAINER_IMAGE?=containerize/${APP}

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/pkg/version.Release=${RELEASE} \
		-X ${PROJECT}/pkg/version.Commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.BuildTime=${BUILD_TIME} " \
		-o ${APP}

test:
	go test -v -race ./...

container: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .
	docker build -t ${CONTAINER_IMAGE}:latest .

run: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		-e "MYSQL_DATABASE=${MYSQL_DATABASE}" \
		-e "TELEGRAM_HORN_URL=${TELEGRAM_HORN_URL}" \
		${CONTAINER_IMAGE}:${RELEASE}

push: container
	docker push ${CONTAINER_IMAGE}:${RELEASE}
	docker push ${CONTAINER_IMAGE}:latest

rsync: 
	rsync -azvh docker-compose.yml root@192.168.20.23:/data/compose/sme-api

gen: 
	xorm reverse mysql root:root@/sme?charset=utf8 pkg/templates/goxorm pkg

