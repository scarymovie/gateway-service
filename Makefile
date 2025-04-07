DOCKER_COMPOSE_FILE := ./docker/docker-compose.yml


.PHONY: up down ps logs build restart migrate-up migrate-down migrate-drop

up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

ps:
	docker compose -f $(DOCKER_COMPOSE_FILE) ps

logs:
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f

build:
	docker compose -f $(DOCKER_COMPOSE_FILE) build

restart: down up