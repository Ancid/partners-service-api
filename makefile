build:
	@docker-compose up --build -d

up:
	@docker-compose up -d

down:
	@docker-compose down

remove:
	@docker-compose down -v

logs:
	@docker logs -f api-service-go_app_1

tests:
	@docker build -t app_test -f Dockerfile.test .
	@docker run app_test

