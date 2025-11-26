FROM golang:1.25.4-trixie

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN go build main.go

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT [ "./main" ]