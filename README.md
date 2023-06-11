#Running postgres on docker

docker run --name workshop-db -e POSTGRES_PASSWORD=goworkshop -p 5432:5432 -d postgres