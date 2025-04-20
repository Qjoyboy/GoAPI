# Makefile для создания миграций
# и применения их к базе данных

DB_DSN := "postgres://postgres:123@localhost:5432/main?sslmode=disable"
MIGRATE := ~/go/bin/migrate -path ./migrations -database $(DB_DSN)


migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go 

gen: 
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
