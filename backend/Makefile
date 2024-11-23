include ../.env

MIGRATION_PATH=db/migrations
DATABASE_URL="postgresql://${DB_USER}:${PASSWORD}@${HOST}:${PORT}/${DB_NAME}?sslmode=disable"

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $$name