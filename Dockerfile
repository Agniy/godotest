FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o aplication main.go

FROM golang:alpine
WORKDIR /go/bin
COPY --from=builder /go/src/app/aplication .
ENTRYPOINT aplication

EXPOSE 8080