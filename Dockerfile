FROM golang:1.21-alpine3.18 AS builder
RUN apk update && apk add --no-cache git
WORKDIR /todos
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /todos
COPY --from=builder /todos/main .
COPY app.env .
COPY db/migration ./db/migration
ADD ./db/migration/000001_init-schema.up.sql /docker-entrypoint-initdb.d

CMD ["/todos/main"]

# FROM golang:alpine
# RUN apk update && apk add --no-cache git
# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build -o main
# ENTRYPOINT [ "app/main" ]
