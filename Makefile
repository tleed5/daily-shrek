.PHONY: build clean deploy deploy-force remove

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/helloShrek main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-force: clean build
	sls deploy --force --verbose

remove:
	sls remove --verbose
