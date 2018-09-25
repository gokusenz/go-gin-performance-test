SERVICE=go-gin-performance-test
REGISTRY=hub.docker.com/u/gokusenz/
COMMIT_SHA=$(shell git rev-parse HEAD)
NAME=$(shell git config user.name)

# Dev commands

test:
	go test -coverprofile=c.out ./... && go tool cover -html=c.out

dep:
	dep ensure -v

dep-update:
	dep ensure -update -v

dev:
	go run main.go

build:
	make dep
	make dep-update
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o entrypoint -a -ldflags '-w -s' main.go

docker:
	docker build -t $(REGISTRY)/$(SERVICE) .
