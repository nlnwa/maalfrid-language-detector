PROJ:=maalfrid
ORG_PATH:=github.com/nlnwa
REPO_PATH:=$(ORG_PATH)/$(PROJ)
VERSION ?= $(shell ./scripts/git-version)

## https://golang.org/cmd/link/
## -w Omit the DWARF symbol table.
## -X Set the value of the string variable in importpath named name to value.
LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

.PHONY: release-binary install-dep api

install:
	@CGO_ENABLED=0 go build -a -tags netgo -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/$(PROJ)

api:
	@$(MAKE) -C ./api

install-dep:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure -vendor-only

release-binary: api install-dep install