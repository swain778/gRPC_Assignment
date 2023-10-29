FROM golang:1.20.5-alpine3.18

WORKDIR /app

COPY . .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]