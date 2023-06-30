FROM alpine:3.9

# The Hugo version
ARG HUGO_VERSION=0.111.3

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz /hugo.tar.gz
RUN tar -zxvf hugo.tar.gz

# We add git to the build stage, because Hugo needs it with --enableGitInfo
RUN apk add --no-cache git

# The source files are copied to /site
COPY bootcamp-content /site
WORKDIR /site

ENTRYPOINT ["../hugo"]
CMD ["serve", "--bind", "0.0.0.0", "--port", "8080", "--verbose", "--verboseLog"]