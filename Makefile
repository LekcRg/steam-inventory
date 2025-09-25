PROJECT_PATH = $(shell pwd)
BIN_FOLDER = $(PROJECT_PATH)/bin
BUILD_PATH := $(BIN_FOLDER)/steam-inventory
LINTER := $(BIN_FOLDER)/golangci-lint
SWAG := $(BIN_FOLDER)/swag
GOOSE := $(BIN_FOLDER)/goose

MAIN_PATH := $(PROJECT_PATH)/cmd/main.go
CFG_YAML := $(PROJECT_PATH)/config.yml

build:
	go build -o $(BUILD_PATH) $(MAIN_PATH)

run:
	go run $(MAIN_PATH) -c=$(CFG_YAML)

lint:
	$(LINTER) run

swag:
	$(SWAG) init -g ./cmd/main.go

betteralign:
	betteralign -apply -test_files ./...

migrate:
	$(GOOSE) up

migrate-down:
	$(GOOSE) down

create-migration:
	$(GOOSE) create $(word 2,$(MAKECMDGOALS)) sql
	@:

install-golangci-lint:
	curl -sSfL \
	  https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh |\
		sh -s -- -b $(BIN_FOLDER) v2.5.0

install-swag:
	GOBIN=$(BIN_FOLDER) go install github.com/swaggo/swag/cmd/swag@v1.16.6

install-goose:
	curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=$(PROJECT_PATH) sh -s v3.25.0

install-betteralign:
	GOBIN=$(BIN_FOLDER) go install github.com/dkorunic/betteralign/cmd/betteralign@v0.7.3

install:
	make install-golangci-lint
	make install-swag
	make install-goose
	make install-betteralign

%:
	@:
