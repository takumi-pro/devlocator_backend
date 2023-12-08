db-connect:
	mysql -u root -P 3307 -p -h 127.0.0.1

# schema spy container
spy-up:
	docker compose -f docker-compose-spy.yml up --build --force-recreate spy
	docker rm spy
	docker compose -f docker-compose-spy.yml up -d --build nginx_schemaspy
spy-down:
	docker compose -f docker-compose-spy.yml down