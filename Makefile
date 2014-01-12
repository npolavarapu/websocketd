# Self contained Go build file that will download and install (locally) the correct
# version of Go, and build our programs. Go does not need to be installed on the
# system (and if it already is, it will be ignored).

# To manually invoke the locally installed Go, use ./go

# Go installation config.
GO_VERSION=1.1.2.linux-amd64
GO_DOWNLOAD_URL=http://go.googlecode.com/files/go$(GO_VERSION).tar.gz

# Build websocketd binary
websocketd: go $(wildcard *.go) $(wildcard libwebsocketd/*.go) go-workspace/src/github.com/npolavarapu/websocketd
	./go get ./go-workspace/src/github.com/npolavarapu/websocketd
	./go fmt github.com/npolavarapu/websocketd/libwebsocketd github.com/npolavarapu/websocketd
	./go build

# Create local go workspace and symlink websocketd into the right location.
go-workspace/src/github.com/joewalnes/websocketd:
	mkdir -p go-workspace/src/github.com/npolavarapu
	ln -s ../../../../ go-workspace/src/github.com/npolavarapu/websocketd

# Setup ./go wrapper to use local GOPATH/GOROOT.
go: go-v$(GO_VERSION)/.done
	@echo '#!/bin/sh' > $@
	@echo mkdir -p $(abspath go-workspace) >> $@
	@echo GOPATH=$(abspath go-workspace) GOROOT=$(abspath go-v$(GO_VERSION)) $(abspath go-v$(GO_VERSION)/bin/go) \$$@ >> $@
	chmod +x $@
	@echo 'Created ./$@ wrapper'

# Download and unpack Go distribution.
go-v$(GO_VERSION)/.done:
	mkdir -p $(dir $@)
	rm -f $@
	@echo Downloading and unpacking Go $(GO_VERSION) to $(dir $@)
	curl --silent --fail $(GO_DOWNLOAD_URL) | tar xzf - --strip-components=1 -C $(dir $@)
	touch $@

# Clean up binary
clean:
	rm -rf websocketd go-workspace
.PHONY: clean

# Also clean up downloaded Go
clobber: clean
	rm -rf go $(wildcard go-v*)
.PHONY: clobber
