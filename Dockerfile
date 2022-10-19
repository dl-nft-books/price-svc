FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/nft-books/price-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/price-svc /go/src/gitlab.com/tokend/nft-books/price-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/price-svc /usr/local/bin/price-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["price-svc"]
