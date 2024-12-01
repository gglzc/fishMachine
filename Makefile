PROTO_DIR := proto
OUT_DIR := .

buildProto:
	protoc -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/game.proto

run:
	go run ./cmd/main.go

dockerup:
	docker compose up 
dockerDown:
	docker compose down -v 

#初始化migrate
firstMigrate:
	migrate create -ext sql -dir migrations -seq create_users_table


migrateUp:
	migrate -database "mysql://fishuser:fishpassword@tcp(127.0.0.1:3306)/fishmachine" -path ./migrations up
migrateDown:
	migrate -database "mysql://fishuser:fishpassword@tcp(127.0.0.1:3306)/fishmachine" -path ./migrations down

help:
	@echo "Available commands:"
	@echo "  make migrate-up          - Run all up migrations"
	@echo "  make migrate-down        - Run all down migrations"
	@echo "  make migrate-force N=X   - Force database version to N"
	@echo "  make migrate-version     - Show current migration version"
	@echo "  make reset-dirty         - Reset dirty state by forcing version"
