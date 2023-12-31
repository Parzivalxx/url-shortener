FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

FROM alpine as production

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

ENTRYPOINT [ "./main" ]
