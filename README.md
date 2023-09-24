# lets-go

### local development (sslmode=disable)

[golang-migrate](https://github.com/golang-migrate/migrate/blob/v4.16.2/database/postgres/TUTORIAL.md)

## create a database with host, password and database name

`psql -h localhost -U postgres -w -c "create database books;"`

`export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/books?sslmode=disable'`

## create sequencial migration up and down file depends on the folder path -dir <> and migration file name -seq <name>

`migrate create -ext sql -dir db/migrations -seq create_users_table`

## run migration

`migrate -database ${POSTGRESQL_URL} -path db/migrations up`
`migrate -database ${POSTGRESQL_URL} -path db/migrations down`
`migrate -database "postgres://postgres:password@localhost:5432/books?sslmode=disable" -path db/migrations up`

## add some data entries into table

`INSERT INTO books (title, author) VALUES ( 'DATA-ORIENTED DESIGN', 'RICHARD FABIAN');`
`INSERT INTO books (title, author) VALUES ( 'Software Engineering at Google', 'Winters');`

### Docker

`docker build -t lets-go-app .`
`docker run -p 8080:8080 lets-go-app`

### Docker compose

## build the services: ensure the app and database services run in a linked environment and can communicate with each other (ensure go application uses `DATABASE_URL` environment to connect to PostgreSQL)

`docker-compose build`

## start the services

`docker-compose up`
