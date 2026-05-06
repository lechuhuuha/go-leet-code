POSTGRES_PORT ?= 55432
ORDERS_ROW_COUNT ?= 20000000
CUSTOMER_COUNT ?= 100000

.PHONY: start stop restart reset logs ps psql explain

start:
	POSTGRES_PORT=$(POSTGRES_PORT) ORDERS_ROW_COUNT=$(ORDERS_ROW_COUNT) CUSTOMER_COUNT=$(CUSTOMER_COUNT) docker compose up --build

stop:
	docker compose down

restart: stop start

reset:
	docker compose down -v
	POSTGRES_PORT=$(POSTGRES_PORT) ORDERS_ROW_COUNT=$(ORDERS_ROW_COUNT) CUSTOMER_COUNT=$(CUSTOMER_COUNT) docker compose up --build

logs:
	docker compose logs -f postgres

ps:
	docker compose ps

psql:
	docker compose exec postgres psql -U query_lab -d query_lab

explain:
	docker compose exec postgres psql -U query_lab -d query_lab -f /queries/explain_current_query.sql
