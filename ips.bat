@echo "Init & Push & Start project"
make vendor
make init
make build-container
make run-container
make push-container
start http://localhost:3008/swagger/