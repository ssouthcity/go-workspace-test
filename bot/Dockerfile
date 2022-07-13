FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o client .

FROM alpine

COPY --from=builder /app/client /bin/client

ENTRYPOINT [ "/bin/client" ]