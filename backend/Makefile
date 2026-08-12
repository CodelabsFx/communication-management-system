APP_NAME := backend
CMD_API := ./cmd/api
CMD_WORKER := ./cmd/worker

.PHONY: all fmt tidy build-api run-api build-worker run-worker

all: tidy build-api

fmt:
	go fmt ./...

tidy:
	"C:\Program Files\Go\bin\go.exe" mod tidy

build-api:
	go build -o bin/$(APP_NAME)-api $(CMD_API)

run-api:
	go run $(CMD_API)

build-worker:
	go build -o bin/$(APP_NAME)-worker $(CMD_WORKER)

run-worker:
	go run $(CMD_WORKER)
