FROM golang:1.10-alpine

RUN apk add --no-cache --update alpine-sdk protobuf protobuf-dev

COPY . /go/src/github.com/nlnwa/maalfrid-language-detector
RUN cd /go/src/github.com/nlnwa/maalfrid-language-detector && make release-binary

FROM scratch
LABEL maintainer="nettarkivet@nb.no"
COPY --from=0 /go/src/github.com/nlnwa/maalfrid-language-detector/maalfrid /
EXPOSE 8672
ENTRYPOINT ["/maalfrid"]
CMD ["serve"]
