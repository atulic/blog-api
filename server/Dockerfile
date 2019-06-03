FROM golang:1.12-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Declare the port on which the webserver will be exposed.
EXPOSE 4000

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# Use CompileDaemon to run the app and rebuild on any file changes.
# We already mounted our working directory in the docker-compose file.
ENTRYPOINT CompileDaemon -log-prefix=false -command="./blog-api"