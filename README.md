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
