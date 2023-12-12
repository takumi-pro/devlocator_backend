db-connect:
	mysql -u root -P 3307 -p -h 127.0.0.1

server-gen:
	oapi-codegen -generate "server" -package openapi  reference/devlocator.yaml > ./gen/server.gen.go

types-gen:
	oapi-codegen -generate "server" -package openapi  reference/devlocator.yaml > ./gen/types.gen.go

# schema spy container
spy-up:
	docker compose -f docker-compose-spy.yml up --build --force-recreate spy
	docker rm spy
	docker compose -f docker-compose-spy.yml up -d --build nginx_schemaspy
spy-down:
	docker compose -f docker-compose-spy.yml down