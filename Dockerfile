FROM golang:alpine

RUN apk add --no-cache --update alpine-sdk protobuf protobuf-dev

COPY . /go/src/github.com/nlnwa/maalfrid-language-detector

RUN cd /go/src/github.com/nlnwa/maalfrid-language-detector \
&& go get github.com/golang/dep/cmd/dep \
&& dep ensure -vendor-only \
&& VERSION=$(scripts/git-version.sh) \
&& CGO_ENABLED=0 \
go install \
-a \
-tags \
netgo -v \
-ldflags "-w -X github.com/nlnwa/maalfrid-language-detector/pkg/version.Version=${VERSION}" \
github.com/nlnwa/maalfrid-language-detector/cmd/...
# -w Omit the DWARF symbol table.
# -X Set the value of the string variable in importpath named name to value.

FROM scratch
LABEL maintainer="nettarkivet@nb.no"

COPY --from=0 /go/bin/maalfrid /

ENV COUNT=5 PORT=8672 MAX_RECV_MSG_SIZE=10000000

ENTRYPOINT ["/maalfrid"]
CMD ["serve"]

EXPOSE 8672
