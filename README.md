# Docker 
## compose up
`docker-compose up --build`

## clean volumes
`docker-compose down -v`

# Migrations
## Create
`migrate create -seq -ext sql -dir .\cmd\migrate\migrations\ create_users` 

## Run
### up
`migrate -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost/social?sslmode=disable" up`
### down
`migrate -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost/social?sslmode=disable" down`

# DOCS
* install swagger following: [this](https://github.com/swaggo/swag)
* follow the instruction in the link above to start
* run `swag init -g ./api/main.go -d cmd,internal && swag fmt`
* open `http://localhost:8080/v1/swagger/index.html` to see the docs