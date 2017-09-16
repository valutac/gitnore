SHELL := /bin/bash

TARGET := gitnore
VERSION := $(shell cat VERSION)

OS := darwin linux windows
ARCH := 386 amd64

build-all: deps
	mkdir -v -p $(CURDIR)/$(VERSION)
	gox -verbose \
		-os "$(OS)" -arch "$(ARCH)" \
