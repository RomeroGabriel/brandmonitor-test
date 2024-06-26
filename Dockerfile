FROM golang:latest AS builder

WORKDIR /app

# Copy Source code 
COPY ./backend .
# Download dependencies
RUN go mod download

# Build the Go binary with necessary compiler flags for optimization
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o brandmonitor-api /app/cmd/main.go

# Use a minimal alpine image for the runtime stage
FROM alpine:latest

WORKDIR /root/

# Remove source code - not necessary anymore
RUN rm -rf /app

# Copy the binary from the builder stage
COPY --from=builder /app/brandmonitor-api .
# COPY .env .

# Expose the application on a specific port
EXPOSE 8080

# Command to run the application
CMD ["./brandmonitor-api"]