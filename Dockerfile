FROM golang:alpine

ENV GO111MODULE=on

RUN apk add --no-cache --update alpine-sdk
WORKDIR /go/src/github.com/nlnwa/maalfrid-language-detector

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN export VERSION=$(./scripts/git-version.sh) \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -v \
-ldflags "-w -s -X github.com/nlnwa/maalfrid-language-detector/pkg/version.Version=${VERSION}" \
./cmd/...

# -w Omit the DWARF symbol table.
# -s Omit symbol table and debug information
# -X Set the value of the string variable in importpath named name to value.


FROM scratch
LABEL maintainer="mariusb.beck@nb.no"

COPY --from=0 /go/bin/maalfrid /

ENV COUNT=5 \
    PORT=8672 \
    MAX_RECV_MSG_SIZE=10000000

ENTRYPOINT ["/maalfrid"]
CMD ["serve"]

EXPOSE 8672
