# Define Shell
SHELL = /bin/sh
PWD = $(shell pwd)
app = pulse
cmd_dir = cmd/pulse
web_dir = web
cmd_out = bin
web_out = bin/template

.PHONY: help
help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Setup dev environment
		yarn --cwd $(PWD)/$(web_dir) install

.PHONY: web
web: ## Build web
		yarn --cwd $(PWD)/$(web_dir) export -o $(PWD)/$(web_out)

# pkger is not working with `pkger -o cmd/pulse` without first running `pkger`.
# So using this weird logic of running it twice
.PHONY: pkger
pkger: web ## Compile web files using pkger
		rm -f pkged.go
		rm -f $(PWD)/$(cmd_dir)/pkged.go
		pkger
		pkger -o $(cmd_dir)/
		rm -f pkged.go

# Helps in testing cmd
.PHONY: dev
dev: web ## Start the development environment
	  yarn --cwd $(PWD)/$(web_dir) dev &
		go run $(PWD)/$(cmd_dir)/

# Remove pkger to stop recompiling of web files
.PHONY: build
build: pkger ## Build pulse command
		go build -o $(PWD)/$(cmd_out)/$(app) $(PWD)/$(cmd_dir)

.PHONY: install
install: pkger ## Build and install pulse command
		go install $(PWD)/$(cmd_dir)

.PHONY: clean
clean: ## Remove all the build files
		rm -r $(PWD)/$(cmd_out)
		rm -r $(PWD)/$(web_out)

