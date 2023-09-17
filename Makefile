db:
	docker run -d --name echo-db -p 5432 -e POSTGRES_PASSWORD=devsample postgres:15.4-alpine3.18 