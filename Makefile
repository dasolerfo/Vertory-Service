postgres: 
	docker run --name Vertory -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Barcelona1234 -d postgres:12-alpine
startDB:
	docker start Vertory
createdb:
	docker exec -it Vertory createdb --username=root --owner=root vertory
dropdb:
	docker exec -it Vertory dropdb vertory
migrateup:
	migrate -path db/migration -database "postgresql://root:Barcelona1234@localhost:5432/vertory?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:Barcelona1234@localhost:5432/vertory?sslmode=disable" -verbose down
sqlc:
	sqlc generate
