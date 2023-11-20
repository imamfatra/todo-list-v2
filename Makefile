DB_URL=postgresql://root:secret@localhost:5432/todo_list?sslmode=disable

network:
	docker network create todos-network

postgres:
	docker run --name pgTodo --network todos-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -ti pgTodo createdb --username=root --owner=root todo_list

dropdb:
	docker exec -ti pgTodo dropdb todo_list

createmigrate:
	migrate create -ext sql -dir db/migration -seq init-schema

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./repository/repository_test.go
	go test -v -cover -short ./service/service_test.go
	go test -v -cover -short ./controller/controller_test.go

server:
	go run main.go

.PHONY: network postgres createdb dropdb createmigrate migrateup migratedown test server