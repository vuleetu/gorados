.PHONY : build install

install : *.go
	@echo "compiling and installing..."
	@go install
	@echo "done"

build : *.go
	go build
