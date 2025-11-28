BIN_DIR := ./build/bin
APP := $(BIN_DIR)/app

MAIN_PKG := ./cmd/app

ENV_FILE := .env


.PHONY: all build clean run

all: build

build:
	@echo "==> Creating build directory..."
	mkdir -p $(BIN_DIR)
	@echo "==> Building $(APP)..."
	go build -o $(APP) $(MAIN_PKG)

clean:
	@echo "==> Cleaning build directory..."
	rm -rf $(BIN_DIR)

run: build
	@echo "==> Running $(APP)..."
	@export $(shell cat $(ENV_FILE) | xargs) && $(APP)
