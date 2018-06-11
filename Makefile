
next:
	docker build --target=build -f ./app/next/Dockerfile ./app

next-prod:
	docker build --target=production -f ./app/next/Dockerfile ./app

test:
	docker-compose \
		-f docker-compose.test.yml \
		-p testing \
			up \
				--build \
				--renew-anon-volumes \
				--exit-code-from test \
				--abort-on-container-exit \
				--remove-orphans
