all: kill-port

package:
	@go get -u github.com/gorilla/mux
	@echo "Package installed"

kill-port:
	@kill -HUP $$(lsof -t -i:8181)
	@echo "Port 8181 is killed"
