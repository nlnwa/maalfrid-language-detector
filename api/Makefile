SUFFIX:= pb.go
DEPS:= $(shell find . -type f -name '*.proto' | sed "s/proto$$/$(SUFFIX)/")
FLAGS:=--go_out=plugins=grpc:.

.PHONY: all install install-deps clean

all: clean install-deps install

install:
	@protoc -I/usr/include -I. $(FLAGS) *.proto

install-deps:
	@go get github.com/golang/protobuf/proto
	@go get github.com/golang/protobuf/protoc-gen-go

clean:
	@$(RM) $(DEPS)
