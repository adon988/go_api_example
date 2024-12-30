swag_init:
	swag init --dir ./cmd/server,./internal/ --output ./docs

server:
	cd cmd/server && go run main.go

automigrate:
	cd cmd/migrate/ && go run main.go -automigrate

migrate:
	cd cmd/migrate && go run main.go -migrate_table=$(migrate_table)

alert_table:
	cd cmd/migrate && go run main.go -alert_table=$(table)
