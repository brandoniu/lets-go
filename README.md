# lets-go

### taskfile

`brew install go-task`

### golangci-lint

`brew install golangci-lint`
`golangci-lint run`
`golangci-lint run --disable-all -E errcheck`

## create a Taskfile.yml

`task assets build`

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

### Kubernetes, running local solution with EKS

- install `kubectl` to configure communicate with cluster
  `brew install kubernetes-cli`
- convert Docker compose to Kubernetes manifests using `kompose`
  `brew install kompose`
- generate Kubernetes manifest for deployments, services, and persistent volume
  `kompose convert -f docker-compose.yml`
- using Helm, a package manger for Kubernetes due to generated manifests not perfect for PostgreSQL deployment
- add the Bitnami helm repository and install PostgreSQL chart
  `helm repo add bitnami https://charts.bitnami.com/bitnami`
  `helm install my-postgres bitnami/postgresql`
- make note of the connection credentials Helm outputs. update app configuration or secrets to use these credentials
- deploy your app
  `kubectl apply -f <generated-deployment-file>.yml`
  `kubectl apply -f <generated-service-file>.yml`
- Ensure the environment variables in your application's deployment manifest (especially DATABASE_URL) are set correctly, pointing to the PostgreSQL service deployed earlier.
- Access the app: If your application's service type is LoadBalancer (common in cloud environments), it will get an external IP address. Use kubectl get svc to check. If you're using Minikube, you can access the app with minikube service <service-name>.

- Storage: Ensure that you're using the right storage solution for your database in Kubernetes. The default storage class might not be suitable for production workloads.

- Configuration & Secrets: Use Kubernetes ConfigMaps and Secrets to manage configurations and sensitive data.

- Scaling: With Kubernetes, you can easily scale your app using ReplicaSets.

- Monitoring & Logging: Consider adding solutions like Prometheus and Grafana for monitoring and Loki or ELK stack for logging.

- Network Policies: Define network policies for better security.
