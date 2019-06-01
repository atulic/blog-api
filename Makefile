test: 
	cd server && go test -v -cover -count=1 -covermode=atomic ./...

start:
	docker-compose up --build

stop:
	docker-compose down

generate-mocks: 
	cd server && mockgen -source=postgres/repository.go -destination mocks/repository.go