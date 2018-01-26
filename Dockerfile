FROM golang:1.9.2-alpine

RUN apk add --no-cache --update alpine-sdk protobuf protobuf-dev

COPY . /go/src/github.com/nlnwa/maalfrid
RUN cd /go/src/github.com/nlnwa/maalfrid && make release-binary

FROM scratch
LABEL maintainer="nettarkivet@nb.no"
COPY --from=0 /go/src/github.com/nlnwa/maalfrid/maalfrid /
EXPOSE 8672
ENTRYPOINT ["/maalfrid"]
CMD ["serve"]
