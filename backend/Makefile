include ../.env

MIGRATION_PATH=db/migrations
DB_URL="postgresql://${DB_USER}:${PASSWORD}@${HOST}:${PORT}/${DB_NAME}?sslmode=disable"

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $$name

.PHONY: migrate-up
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) up

.PHONY: migrate-down
migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) down


.PHONY: migrate-rollback
migrate-rollback:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) down 1

.PHONY: migrate-version
migrate-version:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) version

.PHONY: migrate-drop
migrate-drop:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) drop

.PHONY: dev-setup
dev-setup: migrate-drop migrate-up