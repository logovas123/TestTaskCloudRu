FROM golang:1.23

WORKDIR ${GOPATH}/pvz-service/
COPY . ${GOPATH}/pvz-service/

RUN go build -o /build ./cmd/main.go \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]