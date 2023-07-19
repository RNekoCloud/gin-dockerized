FROM golang:1.20.6-alpine3.18

RUN mkdir /gin-dockerized
WORKDIR /gin-dockerized

COPY go.mod /gin-dockerized

COPY . /gin-dockerized

RUN go mod tidy

RUN go build -o /gin-dockerized/bin/server /gin-dockerized/main.go

EXPOSE 3003

CMD ["/gin-dockerized/bin/server"]