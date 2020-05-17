VERSION=$(shell git describe --always --dirty)
BASENAME=get2pushover
GOOS=$(shell eval $$(go env) && printf "$$GOOS")
GOARCH=$(shell eval $$(go env) && printf "$$GOARCH")
PACKAGE=$(BASENAME)-$(VERSION)-$(GOOS)-$(GOARCH)

all: build

.PHONY: build
build:
	cd src/ && go build -ldflags="-X main.xVersion=$(VERSION)" -o ../build/bin/$(BASENAME)

dist: build
	mkdir -p build/$(BASENAME)
	cp build/bin/$(BASENAME) build/$(BASENAME)/
	cp misc/get2pushover.service build/$(BASENAME)/
	cp misc/config build/$(BASENAME)/
	cd build && tar -czvf $(PACKAGE).tar.gz $(BASENAME)
	rm -rf $(BASENAME)

.PHONY: clean
clean:
	rm -rf ./build
