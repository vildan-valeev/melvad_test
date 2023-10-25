APP_BIN = app/build/app

help:
	@echo "Base commands"
	@echo "\tmake lint\t-> \trunning linter"
	@echo "\tmake local_db_up\t-> \trunning local database"

lint:
	golangci-lint run


local_db_up:
	docker-compose up --build