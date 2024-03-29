FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dl-nft-books/blob-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/blob-svc /go/src/github.com/dl-nft-books/blob-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/blob-svc /usr/local/bin/blob-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blob-svc"]
