FROM golang:1.21-alpine
WORKDIR /app

COPY ./src ./src

WORKDIR /app/src

RUN go mod download


RUN go build -o /app/gin-server /app/src/**.go
WORKDIR /app
RUN rm -rf ./src

EXPOSE 8080
ENTRYPOINT [ "/app/gin-server" ]