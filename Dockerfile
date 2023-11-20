FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
ADD ./db/migration/000001_init-schema.up.sql /docker-entrypoint-initdb.d

EXPOSE 3000
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]