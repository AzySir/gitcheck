.PHONY: *
run:
	go run main.go

build: 
	go build -o gitcheck main.go

deploy: 
	sudo mv gitcheck /usr/local/bin/gitcheck

test: 
	go test -v -cover