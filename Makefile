postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASS=toor -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it postgres12 dropdb simplebank
migrate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

test:
	go test -v -cover ./... 


.PHONY: postgres createdb dropdb migrate migrateup migratedown sqlc
