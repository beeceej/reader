run-backend: build
	docker-compose up reader 

build:
	docker-compose up -d mysql
	docker build -t reader/golang .

.PHONY: run-backend run-common