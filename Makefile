
APP=eth-gas-watcher
BIN="bin"

.PHONY: clean bin build run test

clean:
	@rm -rf $(BIN)/$(APP)
	@rm -rf $(BIN)

bin:
	@mkdir -p $(BIN)

build: bin
	@go build -o $(BIN)/$(APP)

run: build
	@./$(BIN)/$(APP)

test:
	@go run test -v ./...
