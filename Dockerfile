FROM golang:1.17-alpine3.14 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build main.go

FROM golang:1.17-alpine3.14

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE $PORT

CMD ["./main"]