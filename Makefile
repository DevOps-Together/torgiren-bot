all: clean build test

build: 
	go build -o bin/discord-bot .

test:
	go test ./...

clean:
	rm -rf bin

dev:
	bin/discord-bot -log-level=trace -config-file=example/config.yaml
