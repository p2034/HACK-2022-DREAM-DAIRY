psql:
	docker run --name basic-postgres --rm \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=4y7sV96vA9wv46VR \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.1-alpine