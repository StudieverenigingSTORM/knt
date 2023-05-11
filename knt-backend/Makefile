RELEASE_NAME=knt-backend
all: build

docker:
	docker build -t ${RELEASE_NAME} .

dev:
	go run main.go -tags "libsqlite3"

build: ${RELEASE_NAME}

${RELEASE_NAME}: main.go go.mod go.sum internal
	go build -o ./${REALEASE_NAME} .

clean:
	go clean
	rm ${RELEASE_NAME}
