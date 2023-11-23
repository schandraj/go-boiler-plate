# Start from base image
FROM golang:alpine

ARG APPNAME="service-name"

# Set the current working directory inside the container
ADD . $GOPATH/src/"${APPNAME}"
WORKDIR $GOPATH/src/"${APPNAME}"

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy source from current directory to working directory
COPY . .

# Build the application
RUN go build -o service-name .

# Expose necessary port
EXPOSE 8055

# Run the created binary executable after wait for mysql container to be up
CMD ["./service-name", "serve"]