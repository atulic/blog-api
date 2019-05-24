test: 
	go test -v -cover -covermode=atomic ./...

start:
	docker-compose up --build

stop:
	docker-compose down