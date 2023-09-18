postgres:
	docker run --name carRental -p 5432:5432 -e POSTGRES_USER=mikita -e POSTGRES_PASSWORD=mikita1 -d postgres:latest

createdb:
	docker exec -it carRental createdb --username=mikita --owner=mikita car_rental

dropdb:
	docker exec -it carRental dropdb car_rental

migrateup:
	migrate -path db/migrations -database "postgresql://mikita:mikita1@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://mikita:mikita1@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedow