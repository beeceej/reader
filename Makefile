api-up: build mysql-up
	docker-compose run setup
	docker-compose up -d reader

mysql-up:
	docker-compose up -d mysql

down:
	docker-compose run teardown
	docker-compose down

build:
	docker build -t reader/golang .

.PHONY: api-up mysql-up down build
