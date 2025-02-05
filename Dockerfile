FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go
RUN ls -l app
RUN chmod +x app

EXPOSE 8080

CMD ["./app"]