# Define Shell
SHELL = /bin/sh
PWD = $(shell pwd)
app = pulse
cmd_dir = cmd/pulse
frontend_dir = client
cmd_out = bin

.PHONY: help
help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: ## Setup dev environment
		yarn --cwd $(PWD)/$(frontend_dir) install

.PHONY: client
client: dev ## Build client
		yarn --cwd $(PWD)/$(frontend_dir) build
		yarn --cwd $(PWD)/$(frontend_dir) export -o $(PWD)/$(cmd_out)/template

# pkger 
.PHONY: pkger
pkger: client ## Compile client files
	  rm -f $(PWD)/$(cmd_dir)/pkged.go
		pkger
		mv pkged.go $(PWD)/$(cmd_dir)/pkged.go

# Remove pkger to stop recompiling of client files
.PHONY: build
build: pkger ## Build pulse command
		go build -o $(PWD)/$(cmd_out)/$(app) $(PWD)/$(cmd_dir)

.PHONY: install
install: build ## Build and install pulse command
		go install $(PWD)/$(cmd_dir)

.PHONY: clean
clean: ## Remove all the build files
		rm -r $(PWD)/$(cmd_out)

