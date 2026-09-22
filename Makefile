include .env
export

export PROJECT_ROOT=$(shell pwd)


dc-close:
	@docker compose down
env-up:
	@docker compose up -d postgres
env-down:
	@docker compose down postgres
env-port-forward:
	@docker compose up -d port-forwarder
env-port-close:
	@docker compose down port-forwarder
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq is empty. make migrate-create seq=init";\
		exit 1; \
	fi; \
	mkdir -p migrations; \
	docker compose run --rm click-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)" 
migrate-up:
	@make migrate-actions action=up
migrate-down:
	@make migrate-actions action=down
migrate-actions:
	@if [ -z "$(action)" ]; then \
		echo "action is empty. make migrate-actions action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm todo-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"