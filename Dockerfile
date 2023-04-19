# base
FROM golang:1.20.3 as go
WORKDIR /
COPY . .

RUN CGO_ENABLED=0 go build -o /hello hello.go

ENTRYPOINT ["/hello"]
