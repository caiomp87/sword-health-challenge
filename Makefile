migrate-create:
	migrate create -ext sql -dir $$PWD/sql/migrations -seq create_table

migrate-up:
	migrate -database "${DB_URI}" -path sql/migrations up

migrate-down:
	migrate -database "${DB_URI}" -path sql/migrations down

generate-sqlc:
	sqlc generate
