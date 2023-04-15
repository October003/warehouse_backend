run: clean
	@docker-compose up -d --force-recreate --build backend

debug: down
	@docker-compose up -d --force-recreate --build backend

down:
	@docker-compose down

logs:
	@docker logs backend

clean: down
	@sudo rm -rf tmp/logs/*
