# Define Shell
SHELL = /bin/sh
# Opinionated: I like to use absolute path
PWD = $(shell pwd)
app = pulse
pulse_dir = cmd/pulse
explore_dir = cmd/pulseExplore
web_dir = web
cmd_out = bin
web_out = bin/template
miband_dir = tools/miband

.PHONY: help setup web tools grpc dev pkger build install uninstall run clean hard-clean

help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup dev environment
		yarn --cwd $(PWD)/$(web_dir) install
		python3 -m pip install -r $(PWD)/$(miband_dir)/requirements.txt

web: ## Build web
		yarn --cwd $(PWD)/$(web_dir) export -o $(PWD)/$(web_out)

grpc: ## Build proto buffer files
	python3 -m grpc_tools.protoc -Iapi/ \
		--python_out=$(PWD)/$(miband_dir)/src \
		--grpc_python_out=$(PWD)/$(miband_dir)/src \
		api/mibandDevice.proto
	protoc -I api/ \
		-I${GOPATH}/src \
	  --go_out=plugins=grpc:$(PWD)/$(pulse_dir) \
	  api/mibandDevice.proto
	protoc -I api/ \
		-I${GOPATH}/src \
	  --go_out=plugins=grpc:$(PWD)/$(pulse_dir) \
	  api/exploreDevices.proto

tools: ## build the miband library
	cd $(PWD)/$(miband_dir) && python3 setup.py sdist	 # build the package
	cd $(PWD)/$(miband_dir) && python3 setup.py build_ext --inplace # Build .so files
	cd $(PWD)/$(miband_dir) && python3 setup.py bdist_wheel # build .whl file
	#python3 -m pip install /path
	# Add clean instruction too

dev: ## Start the development environment
	  yarn --cwd $(PWD)/$(web_dir) dev &
		go run $(PWD)/$(pulse_dir)/

# pkger is not working with `pkger -o cmd/pulse` without first running `pkger`.
# So using this weird logic of running it twice
pkger: web ## Compile web files using pkger
		rm -f pkged.go
		rm -f $(PWD)/$(pulse_dir)/pkged.go
		pkger
		pkger -o $(pulse_dir)/
		rm -f pkged.go

# Remove pkger to stop recompiling of web files
build: pkger ## Build pulse command
		go build -o $(PWD)/$(cmd_out)/$(app) $(PWD)/$(pulse_dir)

install: ## Build and install pulse command
		go install $(PWD)/$(pulse_dir)

uninstall: ## Uninstall the pulse command and package

run: ## Run the project

clean: ## Remove all the build files
		rm -rf $(PWD)/$(cmd_out)
		rm -rf $(PWD)/$(miband_dir)/mibandPulse.egg-info
		rm -rf $(PWD)/$(miband_dir)/build
		rm -rf $(PWD)/$(miband_dir)/dist

hard-clean: clean ## Remove all the build files and pre-build files (ie. paker.go)
		rm -rf $(PWD)/$(pulse_dir)/pkger.go
		rm -rf $(PWD)/$(web_dir)/node_modules
		rm -rf $(PWD)/$(web_dir)/.next
		rm -rf $(PWD)/$(web_dir)/out
		rm -rf $(PWD)/$(miband_dir)/src/*.c
		rm -rf $(PWD)/$(miband_dir)/src/*.so

