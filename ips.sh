@echo "Init & Push & Start project"
make vendor
make init
make build-container
make run-container
make push-container
xdg-open http://localhost:3008/swagger/
