
#build stage
FROM golang:1.12-rc-alpine3.8 AS builder
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
ENV GO111MODULE=on
COPY . .
RUN apk add bash --no-cache git
RUN go mod download
RUN go build

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates htop
RUN mkdir -p /app
COPY --from=builder /go/src/app /app
WORKDIR /app
LABEL Name=go-health Version=0.0.1
EXPOSE 9000
CMD ["./go-health"]
