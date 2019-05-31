test: 
	go test -cover -count=1 -covermode=atomic ./...

start:
	docker-compose up --build

stop:
	docker-compose down