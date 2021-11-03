all: run

.PHONY: run
run: clean build
	@./bin/server

.PHONY: clean
clean:
	@rm -f  ./bin/server

.PHONY: build
build:
	@go build -o bin/server ./cmd/server