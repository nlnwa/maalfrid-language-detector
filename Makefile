PROJ=maalfrid
ORG_PATH=github.com/nlnwa
REPO_PATH=$(ORG_PATH)/$(PROJ)
VERSION ?= $(shell ./scripts/git-version)

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

.PHONY: release-binary
release-binary:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure -vendor-only
	@CGO_ENABLED=0 go build -a -tags netgo -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/$(PROJ)
