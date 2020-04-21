# Define Shell
SHELL = /bin/sh
cmd_dir = cmd/pulse/main.go
frontend_dir = client
cmd_out = pulse

.PHONY: help
help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: ## Setup dev environment
		yarn --cwd ./$(frontend_dir) install

.PHONY: client
client: dev ## Build client
		yarn --cwd ./$(frontend_dir) build
		yarn --cwd ./$(frontend_dir) export

.PHONY: build
build: client ## Build pulse command
		go build -o $(cmd_out) $(cmd_dir)

.PHONY: clean
clean: ## Remove all the build files
		rm -r $(frontend_dir)/out/
		rm $(cmd_out)


