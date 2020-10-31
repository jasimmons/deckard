BINDIR=bin
BIN_LINUX=deckard-linux-amd64
BIN_MAC=deckard-darwin-amd64
DOCKER_IMAGE_BASE=docker.jasonneeds.coffee/jasimmons/deckard
DOCKER_IMAGE_TAG=$(shell git rev-parse --short HEAD)

.PHONY:
lint:
	golangci-lint run

.PHONY:
test: lint
	go test -v -race -cover ./...

checker/checker.pb.go:
	go generate

identifier/identifier.pb.go:
	go generate

${BINDIR}/${BIN_MAC}: checker/checker.pb.go
	GOOS=darwin GOARCH=amd64 go build -o ${BINDIR}/${BIN_MAC} cmd/server/main.go

${BINDIR}/${BIN_LINUX}: checker/checker.pb.go
	GOOS=linux GOARCH=amd64 go build -o ${BINDIR}/${BIN_LINUX} cmd/server/main.go

.PHONY:
build: ${BINDIR}/${BIN_MAC} ${BINDIR}/${BIN_LINUX}

.PHONY:
docker: build
	docker build -t ${DOCKER_IMAGE_BASE}:${DOCKER_IMAGE_TAG} .

.PHONY:
clean:
	rm -f ${BINDIR}/${BIN_MAC}
	rm -f ${BINDIR}/${BIN_LINUX}
	rm -rf ${BINDIR}
