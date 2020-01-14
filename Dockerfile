FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/dimonchik0036/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/play

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/play /go/bin/play

EXPOSE 7531
ENTRYPOINT ["/go/bin/play"]
