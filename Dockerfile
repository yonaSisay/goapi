# Use the official Golang image as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]
