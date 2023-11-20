FROM postgres:14-alpine3.18
WORKDIR /docker-entrypoint-initdb.d
ADD ./db/todos.sql /docker-entrypoint-initdb.d
EXPOSE 5432