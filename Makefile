.PHONY: run

build: main.go
	docker build -t trivial-server . 

run: build
	docker run --rm -p 8080:8080 trivial-server
