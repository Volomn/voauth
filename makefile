api:
	docker-compose up voauth_api

api_bash:
	docker-compose run voauth_api bash

build:
	docker-compose build

down:
	docker-compose down

test:
	cd backend && go test -cover ./...
