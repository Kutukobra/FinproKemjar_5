FROM golang:1.25.4-trixie

WORKDIR /app

COPY . .

RUN go mod download

RUN go build main.go

ENV AppPort=8888

EXPOSE 8888

ENTRYPOINT [ "./main" ]