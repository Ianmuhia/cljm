build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose down

dev:
	go run ./cmd/ *.go

reset:
	docker-compose down --remove-orphans --volumes 




#check for pg_admin ip 
#docker inspect <postgre_container_id> | grep IPAddress