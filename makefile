api:
	docker-compose up voauth_api

api_bash:
	docker-compose run voauth_api bash

build:
	docker-compose build

down:
	docker-compose down

ecr_login:
	aws ecr get-login-password --region eu-west-1 --profile shaibu_volomn | docker login --username AWS --password-stdin 206722258093.dkr.ecr.eu-west-1.amazonaws.com

build_api_prod_image:
	cd backend && docker build -t 206722258093.dkr.ecr.eu-west-1.amazonaws.com/voauth_api:prod -f Dockerfile .

build_web_prod_image:
	cd frontend && docker build --build-arg API_BASE_URL=https://api.voauth.volomn.io/api -t 206722258093.dkr.ecr.eu-west-1.amazonaws.com/voauth_web:prod -f Dockerfile .

push_api_prod_image:
	docker push 206722258093.dkr.ecr.eu-west-1.amazonaws.com/voauth_api:prod

push_web_prod_image:
	docker push 206722258093.dkr.ecr.eu-west-1.amazonaws.com/voauth_web:prod

build_and_push_web_to_prod: ecr_login build_web_prod_image push_web_prod_image

build_and_push_api_to_prod: ecr_login build_api_prod_image push_api_prod_image

build_and_push_images_to_prod: build_and_push_web_to_prod build_and_push_api_to_prod

test:
	cd backend && go test -cover ./...
