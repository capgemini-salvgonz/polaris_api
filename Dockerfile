# Use the official Go image as the base image
FROM golang

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o api-app .

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./api-app"]
