DOCKER_COMPOSE = docker compose -f testbackends/docker-compose.yaml

# запуск сервиса
.PHONY: up
up:
	@echo "Запуск сервиса..."
	@docker compose up --build

# завершение работы сервиса
.PHONY: down
down:
	@echo "Остановка сервиса..."
	@docker compose down -v	


# перезапуск сервиса
.PHONY: restart
restart: down up

# очистка docker
.PHONY: clean
clean:
	@echo "Очистка Docker..."
	@docker system prune -af