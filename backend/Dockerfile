FROM golang:1.25.4-trixie AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install templ CLI for generating template files
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy all source code
COPY . .

# Generate templ templates
RUN templ generate

# Build the application
RUN go build -o main .

# Final stage - smaller image
FROM debian:bookworm-slim

WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/main .

# Copy static files and templates
COPY --from=builder /app/app/static ./app/static
COPY --from=builder /app/app/template ./app/template

ENV PORT=8080
ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["./main"]