# Load .env if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif


migrate:
	goose -dir db/sql/schema postgres "$(DB_URL)" up

migrate-down:
	goose -dir db/sql/schema postgres "$(DB_URL)" down

migrate-status:
	goose -dir db/sql/schema postgres "$(DB_URL)" status
