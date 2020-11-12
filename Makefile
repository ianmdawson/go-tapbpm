build: test
	go build

build-run: build
	./go-tapbpm

install: test
	go install

run:
	go run main.go taptracker.go

test:
	go test -v
