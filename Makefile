
frontend:
	docker build --target=build -f ./app/frontend/Dockerfile ./app

frontend-prod:
	docker build --target=production -f ./app/frontend/Dockerfile ./app

frontend-test:
	docker-compose \
		-f docker-compose.test.yml \
		-p testing \
			up \
				--build \
				--renew-anon-volumes \
				--exit-code-from test \
				--abort-on-container-exit \
				--remove-orphans

test: frontend-test

build: frontend

production: frontend-prod

.PHONY: build
