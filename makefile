# Makefile for the web_app project
# Author: John Doe
# Date: 2026-03-16
# Version: 1.0.0
# Description: This Makefile is used to build and run the web_app project
# Usage: make dev - to run the project in development mode
#        make start - to run the project in production mode
#        make build - to build the project
#        make clean - to clean the project

# Variables
APP_NAME = webapp
APP_VERSION = 1.0.0
APP_AUTHOR = Mark Wayne Menorca
APP_DATE = 2026-03-16
APP_DESCRIPTION = This is a web application built with Echo and MongoDB
APP_LICENSE = MIT
APP_LICENSE_URL = https://opensource.org/licenses/MIT
APP_LICENSE_TEXT = MIT License

APP_DIST_DIR = dist
APP_BUILD_DIR = build
APP_SRC_DIR = app
APP_MAIN_FILE = main.go
APP_CONFIG_DIR = config
APP_LOG_DIR = log
APP_TEMP_DIR = temp
APP_DATA_DIR = data
APP_CACHE_DIR = cache
APP_TEMPLATES_DIR = templates
APP_VIEWS_DIR = views

# Platform / architecture variables (override when calling make)
GOOS  ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Platform-specific executable suffix
ifeq ($(GOOS),windows)
	APP_EXE_SUFFIX = .exe
else
	APP_EXE_SUFFIX =
endif

APP_BIN = $(APP_NAME)-$(GOOS)-$(GOARCH)$(APP_EXE_SUFFIX)


.PHONY: dev start build clean test

dev: 
	air

start:
	go run $(APP_SRC_DIR)/$(APP_MAIN_FILE)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_BUILD_DIR)/$(APP_BIN) $(APP_SRC_DIR)/$(APP_MAIN_FILE)

test:
	go test ./tests/...

clean:
	rm -Rf $(APP_BUILD_DIR)/$(APP_BIN)
	rmdir $(APP_BUILD_DIR)