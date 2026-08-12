FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go mod download
COPY . ./
RUN go build -o /app/backend ./cmd/api

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/backend /usr/local/bin/backend
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/backend"]
