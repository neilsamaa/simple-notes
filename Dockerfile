# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Golang binary
RUN go build -o aplikasi-notes-go .

# Set the startup command to run the Golang binary
CMD ["./aplikasi-notes-go"]