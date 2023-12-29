# build stage
FROM golang:alpine AS builder
ADD . /go/src/github.com/feiskyer/openai-copilot
RUN cd /go/src/github.com/feiskyer/openai-copilot && \
    apk update && apk add --no-cache gcc musl-dev openssl && \
    CGO_ENABLED=0 go build -o _out/openai-copilot ./cmd/openai-copilot

# Final image
FROM alpine
# EXPOSE 80
WORKDIR /

RUN apk add --update curl wget python3 py3-pip && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /etc/openai-copilot

COPY --from=builder /go/src/github.com/feiskyer/openai-copilot/_out/openai-copilot /usr/local/bin/

USER copilot
ENTRYPOINT [ "/usr/local/bin/openai-copilot" ]
