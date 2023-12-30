# GCP setting
IMAGE=asia-northeast1-docker.pkg.dev/sigma-method-409207/devlocator-app/devlocator

image-build:
	docker image build -t $(IMAGE):latest --target production --platform linux/amd64 -f docker/golang/Dockerfile .
	
add-tag:
	docker image tag devlocator $(IMAGE):latest

image-push:
	docker image push $(IMAGE):latest

deploy:
	gcloud run deploy devlocator \
	--image $(IMAGE) \
	--port 8000 \
	--platform=managed \
	--allow-unauthenticated \
	--region asia-northeast1

# docker
up:
	docker compose up -d

down:
	docker compose down

destroy:
	docker compose down --rmi all --volumes

db-connect:
	mysql -u takumi -P 3307 -p -h 127.0.0.1 devlocator

server-gen:
	oapi-codegen -generate "server" -package openapi  reference/devlocator.yaml > ./openapi/server.gen.go

types-gen:
	oapi-codegen -generate "types" -package openapi  reference/devlocator.yaml > ./openapi/types.gen.go

prod-build:
	docker build -t devlocator:latest --target production --platform linux/amd64 -f docker/golang/Dockerfile .

prod-run:
	docker container run --name devlocator devlocator

# schema spy container
spy-up:
	docker compose -f docker-compose-spy.yml up --build --force-recreate spy
	docker rm spy
	docker compose -f docker-compose-spy.yml up -d --build nginx_schemaspy
spy-down:
	docker compose -f docker-compose-spy.yml down
