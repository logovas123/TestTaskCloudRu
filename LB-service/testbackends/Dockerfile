FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /server .


EXPOSE 8088

CMD ["/server"]
