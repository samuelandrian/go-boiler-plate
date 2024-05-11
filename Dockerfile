# Use the official Golang image to build the application
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy
RUN go mod vendor

# Copy the source code into the container
COPY . .


# Build the Go app
RUN CGO_ENABLED=0 GO111MODULE=on go build -o myapp /app/cmd/main.go

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /app


#set local timezone jakarta
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" >  /etc/timezone


# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/myapp .

# recopy so files will exist in container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080


# Command to run the executable with argument -env
CMD ["sh", "-c", "./myapp -env $APP_ENV"]

# in shell run this 2 command
# docker build -t go-boiler-plate:latest .
# docker run -p 8080:8080 --env-file .env go-boiler-plate