test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out ./...

testcover:	## Open coverage file as html
	go tool cover -html=c.out

build: 		## Build executable
	set	GOOS=$(os)& set	GOARCH=$(arch)& go build -o ./$(name) main.go

buildimage:	## Build docker image
	docker build -t financial-ledger:$(profile) -f ./docker/Dockerfile-$(profile) .

runimage:	## Run docker image
	docker run -d --name ledger --rm -p 8080:8080 financial-ledger:$(profile)

fmt:		## Formats code and then swagger annotations
	go fmt ./...
	swag fmt