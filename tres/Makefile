BINARY := chronon
DOCKERVER :=`cat VERSION`
.DEFAULT_GOAL := linux

linux:
	cd cmd/$(BINARY) ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
	docker build  --tag="earthcube/chronon:$(DOCKERVER)"  --file=./build/Dockerfile .

dockerlatest:
	docker build  --tag="earthcube/chronon:latest"  --file=./build/Dockerfile .
