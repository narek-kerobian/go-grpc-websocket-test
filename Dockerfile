# Compile application binary
FROM golang:1.21-bookworm AS build-stage

WORKDIR /opt/go-grpc-websocket-test

COPY . /opt/go-grpc-websocket-test

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build

# Run the application as a service
FROM debian:bookworm

RUN set -e && apt-get update && apt-get install -y --no-install-recommends \
    caddy \
    supervisor

WORKDIR /

# Copy compiled binary
COPY --from=build-stage /opt/go-grpc-websocket-test/go-grpc-websocket-test /usr/bin/go-grpc-websocket-test

# Copy configs
COPY docker/supervisord.conf /etc/supervisor/supervisord.conf
COPY docker/Caddyfile /etc/Caddyfile
COPY docker/entrypoint.sh /root/entrypoint.sh

# Set exposed ports
EXPOSE ${APP_PORT}

# Define the entrypoint
ENTRYPOINT ["sh", "/root/entrypoint.sh"]

