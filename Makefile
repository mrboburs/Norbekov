
run-go:
	go run command/main.go

run-psql:
	sudo docker start mediumuz

run-redis:
	sudo docker start redisdb

start-psql:
	sudo docker run medium-db 

start-redis:
	sudo docker run redis-test-instance 

swag:
	swag init -g command/main.go

migrate-up:
	migrate -path ./schema -database 'postgresql://postgres:1996@localhost:5436/mediumuz?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgresql://postgres:1996@localhost:5434/mediumuz?sslmode=disable' down




