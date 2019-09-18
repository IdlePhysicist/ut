OS=darwin
CGO=0

BUILD=build
INSTALL=$(GOPATH)/bin

default: build

build: clean
	env CGO_ENABLED=$(CGO) GOOS=$(OS) go build -o $(BUILD)/ut main.go

clean: 
	rm -f $(BUILD)/*
	touch $(BUILD)/.keep

install: build
	mv $(BUILD)/ut $(INSTALL)
