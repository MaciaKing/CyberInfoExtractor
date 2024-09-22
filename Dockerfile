# Use the official Go image to build the application
FROM golang:1.22.5

# Set the Current Working Directory inside the container
WORKDIR /go/src/app/

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download -x

# Install compile daemon for hot reloading
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# Expose port 80 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT CompileDaemon --build="go build /go/src/app/cmd/main/main.go" --command="./main"
