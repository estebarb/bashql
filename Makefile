all: format install

format:
	go fmt common/*.go
	go fmt bqlselect/*.go
	go fmt bqlfrom/*.go
	go fmt bqlwhere/*.go
	go fmt bqlwhenever/*.go
	go fmt bqljoin/*.go
	
install:
	go install github.com/estebarb/bashql/...
