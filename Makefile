add :
		cd api && go install

install :
		docker exec client yarn install

rebuild:
		docker exec client npm rebuild node-sass
		docker restart client

run :
		mkdir -p docker/data
		docker-compose up --build -d
		cd api && go run main.go

build :
		cd api && go build main.go

run-build :
		docker-compose up --build -d
		./api/main

down :
		docker-compose down
