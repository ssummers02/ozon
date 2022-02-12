build:
	docker-compose build link-app

run:
	docker-compose up link-app

test:
	go test -v ./...

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go

