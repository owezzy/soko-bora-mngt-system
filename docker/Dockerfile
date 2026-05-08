FROM golang:1.24-alpine AS builder

WORKDIR /mallbots

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -ldflags="-s -w" -v -o monolith ./cmd/mallbots

FROM alpine:latest AS runtime

COPY --from=builder /mallbots/docker/wait-for ./wait-for
RUN chmod +x ./wait-for

# Copy the binary to the production image from the builder stage.
COPY --from=builder /mallbots/monolith ./monolith

CMD ["./monolith"]