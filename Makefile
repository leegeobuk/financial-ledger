test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out . ./api ./cfg ./db

testcover:	## Open coverage file as html
	go tool cover -html=c.out

build: 		## Build executable
	set	GOOS=$(os)& set	GOARCH=$(arch)& go build -o ./$(name) main.go

buildimage:	## Build docker image
	docker build -t household-ledger:$(profile) -f ./docker/Dockerfile-$(profile) .

runimage:	## Run app docker image
	docker run -d --rm --network household-ledger --name ledger -p 8080:8080 household-ledger:$(profile)

fmt:		## Formats code and swagger annotations
	go fmt ./...
	swag fmt

setup:		## Setup dependencies using docker-compose
	docker-compose -f ./docker/docker-compose.yaml --env-file ./docker/.env up -d

teardown:	## Tear down dependencies using docker-compose
	docker-compose -f ./docker/docker-compose.yaml down