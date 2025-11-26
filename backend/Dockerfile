FROM golang:1.25.4-trixie

WORKDIR /app

COPY . .

RUN go mod download

RUN go build main.go

EXPOSE 8080

ENTRYPOINT [ "./main" ]