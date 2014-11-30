all: format install

format:
	go fmt common/*.go
	go fmt bqlselect/*.go
	go fmt bqlfrom/*.go
	go fmt bqlwhere/*.go
	go fmt bqlwhenever/*.go
	go fmt bqljoin/*.go
	go fmt bqlsum/*.go
	go fmt bqlcount/*.go
	go fmt bqlavg/*.go
	go fmt bqlmax/*.go
	go fmt bqlmin/*.go
	go fmt bqldistinct/*.go
	
install:
	go install github.com/estebarb/bashql/...
