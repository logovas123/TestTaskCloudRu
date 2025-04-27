FROM golang:1.22

WORKDIR ${GOPATH}/lb-service/
COPY . ${GOPATH}/lb-service/

RUN go build -o /build ./cmd/loadbalancer/main.go \
    && go clean -cache -modcache

EXPOSE 8080