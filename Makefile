build:
	docker-compose build

up:
	docker-compose up

dev:
	go run ./cmd/ *.go

reset:
	docker-compose down --remove-orphans --volumes 




#check for pg_admin ip 
#docker inspect <container id> | grep IPAddress