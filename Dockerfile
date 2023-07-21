### BUILDER ###
FROM golang:1.20.6-alpine as builder

# ENV GO111MODULE=on

# Install git.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download 

# Copy app
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

### FINAL ###
FROM golang:1.20.6-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/views ./views
COPY --from=builder /app/public ./public  

# Run the executable
CMD ["./main"]