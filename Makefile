POSTGRES_HOST=host
POSTGRES_PORT=prot
POSTGRES_DATABASE=database
POSTGRES_USER=user
POSTGRES_PASSWORD=password

-include .env

DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

print:
	echo "$(DB_URL)"

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

migrateup:
		migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
		migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY:	start migrateup migratedown