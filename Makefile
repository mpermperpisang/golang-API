all: kill-port run

package:
	@go get -u github.com/gorilla/mux
	@go get -u github.com/nwidger/jsoncolor
	@echo "Package installed"

kill-port:
	@kill -9 $$(lsof -t -i:8181)
	@echo "Port 8181 is killed"

run:
	@echo "Running code"
	go run main.go
