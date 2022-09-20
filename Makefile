all: clean lint build test test-cover

build: 
	go build -o bin/discord-bot .

test:
	go test ./...

test-cover:
	go test -cover ./...
clean:
	rm -rf bin

dev:
	bin/discord-bot -log-level=trace -config-file=example/config.yaml

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run -v

docker:
	docker build -t torgiren-bot:local .
