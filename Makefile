DOCKER_COMPOSE_RUN_WITH_SERVERS = docker compose -f LB-service/testbackends/docker-compose.yaml
DOCKER_COMPOSE_LB_SERVICE = docker compose -f LB-service/docker-compose.yaml	

# запуск сервиса
.PHONY: up-lb
up-lb:
	@echo "Запуск сервиса Load Balancer..."
	@$(DOCKER_COMPOSE_LB_SERVICE) --env-file ./LB-service/.env up --build

# завершение работы сервиса
.PHONY: down-lb
down-lb:
	@echo "Остановка сервиса Load Balancer..."
	@$(DOCKER_COMPOSE_LB_SERVICE) --env-file ./LB-service/.env down -v	


# перезапуск сервиса
.PHONY: restart-lb
restart-lb: down-lb up-lb

# запуск сервиса вместе с тестовыми серверами
.PHONY: run-lb-with-servers
run-lb-with-servers:
	@echo "Запуск сервиса вместе с тестовыми серверами..."
	@$(DOCKER_COMPOSE_RUN_WITH_SERVERS) --env-file ./LB-service/.env down -v
	@$(DOCKER_COMPOSE_RUN_WITH_SERVERS) --env-file ./LB-service/.env up --build

# очистка docker
.PHONY: clean
clean:
	@echo "Очистка Docker..."
	@docker system prune -af
