FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o server .

FROM alpine

COPY --from=builder /app/server /bin/server

EXPOSE 5000

ENTRYPOINT [ "/bin/server" ]
