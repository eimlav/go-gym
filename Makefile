build:
	go build;

run:
	./go-gym

reset-db:
	rm go-gym.db

test:
	go test ./...