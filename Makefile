test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out . ./api ./db

testcover:	## Open coverage file as html
	go tool cover -html=c.out

build: 		## Build executable
	set	GOOS=$(os)& set	GOARCH=$(arch)& go build -o ./$(name) main.go

buildimage:	## Build docker image
	docker build -t financial-ledger:$(profile) -f ./docker/Dockerfile-$(profile) .

runimage:	## Run app docker image
	docker run -d --name ledger --rm -p 8080:8080 financial-ledger:$(profile)

fmt:		## Formats code and then swagger annotations
	go fmt ./...
	swag fmt

setup:		## Setup dependencies using docker-compose
	docker-compose -f ./docker/docker-compose.yaml --env-file ./docker/.env up -d

teardown:	## Tear down dependencies using docker-compose
	docker-compose -f ./docker/docker-compose.yaml down