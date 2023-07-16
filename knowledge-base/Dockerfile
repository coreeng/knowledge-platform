# Stage 1: Build and test stage
FROM golang:1.18-alpine AS build

# Set the working directory in the container
WORKDIR /app

# Copy the source code to the container
COPY content/ /content
COPY structure/ /structure

WORKDIR /structure
# Download the Go modules
RUN go mod download
RUN apk add --no-cache gcc musl-dev

# Build the Go program
RUN go build -o weight_generator

# Run the Go tests
RUN go test -v ./...

# Generate weights using the Go program
RUN ./weight_generator

FROM alpine:3.9

# The Hugo version
ARG HUGO_VERSION=0.111.3

# Add a non-root user
RUN addgroup -S hugo && adduser -S -G hugo hugo

# Download and extract Hugo
ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz /hugo.tar.gz
RUN tar -zxvf hugo.tar.gz && \
    rm -f hugo.tar.gz && \
    mv hugo /usr/local/bin/hugo && \
    chmod +x /usr/local/bin/hugo

# We add git to the build stage, because Hugo needs it with --enableGitInfo
RUN apk add --no-cache git

# Set the working directory
WORKDIR /site

# Copy the source files
COPY . .
# Copy the built Go program from the build stage
COPY --from=build /structure/weight_generator structure/
# Set ownership and permissions
RUN chown -R hugo:hugo /site

# Set the non-root user
USER hugo
RUN ./structure/weight_generator
# Run Hugo
ENTRYPOINT ["hugo"]
CMD ["serve", "--bind", "0.0.0.0", "--port", "8080", "--verbose", "--verboseLog"]