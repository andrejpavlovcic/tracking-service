createdb:
	docker exec -it postgres createdb --username=admin --owner=admin account

dropdb:
	docker exec -it postgres dropdb --username=admin account 

migrateup:
	migrate -path internal/db/migrations -database "postgresql://admin:admin@localhost:5432/account?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgresql://admin:admin@localhost:5432/account?sslmode=disable" -verbose down

test:
	docker-compose up -d
	docker exec -it tracking_system-postgres-1 createdb --username=admin --owner=admin test
	migrate -path internal/db/migrations -database "postgresql://admin:admin@localhost:5432/test?sslmode=disable" -verbose up
	docker-compose exec -e DB_NAME=test -T server go test ./...
	docker exec -it tracking_system-postgres-1 dropdb --username=admin test
	docker-compose down

client:
	docker attach client

load_test:
	go run tests/main.go

.PHONY: createdb dropdb migrateup migratedown