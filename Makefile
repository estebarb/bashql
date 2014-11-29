all: format install

format:
	go fmt common/*.go
	go fmt bqlselect/*.go
	go fmt bqlfrom/*.go
	
install:
	go install github.com/estebarb/bashql/...
