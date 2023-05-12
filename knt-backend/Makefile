RELEASE_NAME=knt-backend
all: build

docker:
	docker build -t ${RELEASE_NAME}:$(shell date +%Y%m%d) -t ${RELEASE_NAME}:latest  .

docker-up:
	docker compose up -d --force-recreate

docker-down:
	docker compose down

docker-test: docker
	docker compose up --force-recreate

dev:
	go run main.go -tags "libsqlite3"

build: ${RELEASE_NAME}

${RELEASE_NAME}: main.go go.mod go.sum internal
	go build -o ./${RELEASE_NAME} .

clean:
	go clean
