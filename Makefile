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
		go get ../...

.PHONY: web
web: ## Build web
		yarn --cwd $(PWD)/$(web_dir) export -o $(PWD)/$(web_out)

.PHONY: lib
lib: ## build the miband library
	python setup.py sdist	 # build the package
	python setup.py build_ext --inplace # Build .so files
	python setup.py bdist_wheel # build .whl file
	python3 -m pip install /path
	# Add clean instruction too

# Helps in testing cmd
.PHONY: dev
dev: ## Start the development environment
	  yarn --cwd $(PWD)/$(web_dir) dev &
		go run $(PWD)/$(cmd_dir)/

# pkger is not working with `pkger -o cmd/pulse` without first running `pkger`.
# So using this weird logic of running it twice
.PHONY: pkger
pkger: web ## Compile web files using pkger
		rm -f pkged.go
		rm -f $(PWD)/$(cmd_dir)/pkged.go
		pkger
		pkger -o $(cmd_dir)/
		rm -f pkged.go

# Remove pkger to stop recompiling of web files
.PHONY: build
build: pkger ## Build pulse command
		go build -o $(PWD)/$(cmd_out)/$(app) $(PWD)/$(cmd_dir)

.PHONY: install
install: ## Build and install pulse command
		go install $(PWD)/$(cmd_dir)

.PHONY: uninstall
uninstall: ## Uninstall the pulse command and package

.PHONY: run
run: ## Run the project

.PHONY: clean
clean: ## Remove all the build files
		rm -r $(PWD)/$(cmd_out)
		rm -r $(PWD)/$(web_out)

