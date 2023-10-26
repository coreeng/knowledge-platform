FROM golang:1.21-alpine3.18

# Install kubectl
RUN apk update &&  \
    apk add --no-cache build-base=0.5-r3 &&  \
    apk add --no-cache curl=8.4.0-r0 && \
    apk add --no-cache openssh=~9.3 && \
    curl -LO "https://dl.k8s.io/release/v1.26.1/bin/linux/amd64/kubectl"
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl


#Install helm
ENV BASE_URL="https://get.helm.sh"
ENV HELM_VERSION="3.13.1"
RUN case `uname -m` in \
        x86_64) ARCH=amd64; ;; \
        armv7l) ARCH=arm; ;; \
        aarch64) ARCH=arm64; ;; \
        ppc64le) ARCH=ppc64le; ;; \
        s390x) ARCH=s390x; ;; \
        *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac && \
    apk add --update --no-cache wget git bash yq && \
    wget ${BASE_URL}/helm-v${HELM_VERSION}-linux-${ARCH}.tar.gz -O - | tar -xz && \
    mv linux-${ARCH}/helm /usr/bin/helm && \
    chmod +x /usr/bin/helm && \
    rm -rf linux-${ARCH} \

WORKDIR multi-tenancy/

# Install cuelang
RUN go install cuelang.org/go/cmd/cue@v0.6.0

COPY . ./

