# Docker-bin-go: Go Dev in Docker

1. `chmod +x ./bin/go`
2. `docker compose up --build`
3. `./bin/go mod init %module%`
4. `make tidy`
5. `docker compose down`
6. `docker compose up`
7. `make run`, `make test`, `make lint`