psql:
	# https://github.com/sameersbn/docker-postgresql/issues/112
	docker run --name basic-postgres --rm \
	-e POSTGRES_DB=$DATABASE_NAME \
	-e POSTGRES_USER=$DATABASE_USER \
	-e POSTGRES_PASSWORD=$DATABASE_PASSWORD \
	-p 5432:5432 -it postgres

build:
	go build -o bin/server .

start-auth: build
	./bin/server auth

start-api: build
	./bin/server api