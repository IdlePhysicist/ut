os=darwin
cgo=0
version="0.1"
ldflags="-X main.commit=`git rev-list -1 HEAD | head -c 8` -X main.version=$(version)"

BUILD=build
CMD=cmd
INSTALL=$(GOPATH)/bin

default: clean build

build:
	env CGO_ENABLED=$(cgo) GOOS=$(os) go build --ldflags $(ldflags) -o $(BUILD)/ut $(CMD)/main.go

clean: 
	rm -f $(BUILD)/*
	touch $(BUILD)/.keep

install: build
	mv $(BUILD)/ut $(INSTALL)

uninstall:
	rm $(INSTALL)/ut
