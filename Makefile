APP_NAME=pong
CMD=.
BUILD_DIR=build

.PHONY: clean run

clean:
	rm -rf $(BUILD_DIR)

run:
	go run $(CMD)/main.go

build:
	go build -o $(BUILD_DIR)/flappy.exe .