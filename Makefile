build: test
	go build

build-run: build
	./go-tapbpm

run:
	go run main.go taptracker.go

test:
	go test *_test.go

