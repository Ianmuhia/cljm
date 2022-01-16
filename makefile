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

shell:
	docker exec -it clj-dev bash && cd

pgip:
	docker inspect postgres_clj | grep IPAddress

pgdrop:
	docker stop postgres_clj && docker rm postgres_clj


start:
	air

