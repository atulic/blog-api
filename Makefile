test: 
	go test -v -cover -count=1 -covermode=atomic ./...

start:
	docker-compose up --build

stop:
	docker-compose down