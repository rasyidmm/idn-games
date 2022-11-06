setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/client.yml.example internal/config/client/client.yml
	@cp internal/config/example/rest.yml.example internal/config/server/rest.yml
	@echo " --- Done Setup and generate configuration --- "

rest:
	@go run main.go

build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main main.go