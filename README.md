Information.

for local test see ---> config/config.yaml 

docker pull postgres:latest

docker run
    --name myPostgresDb
    -p 32768:5432
    -e POSTGRES_USER=postgres
    -e POSTGRES_PASSWORD=postgrespw
    -e POSTGRES_DB=deepflower
    -d postgres

go run cmd/deepflower/main.go