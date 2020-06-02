FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY ./server/main.go .

RUN CGO_ENABLED=0 go build -o /go/bin/play

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/play /go/bin/play

EXPOSE 7531
ENTRYPOINT ["/go/bin/play"]
