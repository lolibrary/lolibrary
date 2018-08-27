
frontend:
	docker build --target=build -f ./app/frontend/Dockerfile ./app

frontend-prod:
	docker build --target=production -f ./app/frontend/Dockerfile ./app

frontend-test:
	docker-compose -p frontend-test -f docker-compose.test.yml build frontend && \
	docker-compose -p frontend-test -f docker-compose.test.yml run --rm frontend

api-test:
	CACHE_DRIVER=array QUEUE_DRIVER=sync docker-compose -p api-test -f docker-compose.test.yml build api && \
	docker-compose -p api-test -f docker-compose.test.yml run --rm api

test: frontend-test api-test

build: frontend

production: frontend-prod

.PHONY: build
