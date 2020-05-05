# Define Shell
SHELL = /bin/sh
PWD = $(shell pwd)
app = pulse
cmd_dir = cmd/pulse
web_dir = web
cmd_out = bin
web_out = bin/template

.PHONY: help setup web tools grpc dev pkger 

help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup dev environment
		yarn --cwd $(PWD)/$(web_dir) install
		go get ../...

web: ## Build web
		yarn --cwd $(PWD)/$(web_dir) export -o $(PWD)/$(web_out)

tools: ## build the miband library
	python setup.py sdist	 # build the package
	python setup.py build_ext --inplace # Build .so files
	python setup.py bdist_wheel # build .whl file
	python3 -m pip install /path
	# Add clean instruction too

grpc: ## Build proto buffer files
	protoc -I api/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:api \
		api/api.proto
	protoc -I api/ \
		--python_out=$(PWD)/tools/miband/src 

dev: ## Start the development environment
	  yarn --cwd $(PWD)/$(web_dir) dev &
		go run $(PWD)/$(cmd_dir)/

# pkger is not working with `pkger -o cmd/pulse` without first running `pkger`.
# So using this weird logic of running it twice
pkger: web ## Compile web files using pkger
		rm -f pkged.go
		rm -f $(PWD)/$(cmd_dir)/pkged.go
		pkger
		pkger -o $(cmd_dir)/
		rm -f pkged.go

# Remove pkger to stop recompiling of web files
build: pkger ## Build pulse command
		go build -o $(PWD)/$(cmd_out)/$(app) $(PWD)/$(cmd_dir)

install: ## Build and install pulse command
		go install $(PWD)/$(cmd_dir)

uninstall: ## Uninstall the pulse command and package

run: ## Run the project

clean: ## Remove all the build files
		rm -r $(PWD)/$(cmd_out)
		rm -r $(PWD)/$(web_out)

