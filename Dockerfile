FROM golang:1.23 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY cmd cmd
COPY pkg pkg

# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o /prism-proxy ./

FROM alpine:3.21

WORKDIR /app

COPY --from=build /prism-proxy /prism-proxy

ENTRYPOINT ["/prism-proxy"]
CMD ["-config", "/app/config.yaml"]
