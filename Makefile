migrate-create:
	migrate create -ext sql -dir $$PWD/db/migrations -seq ${NAME}

migrate-up:
	migrate -database "${DB_URI}" -path db/migrations up

migrate-down:
	migrate -database "${DB_URI}" -path db/migrations down

generate-sqlc:
	sqlc generate

run:
	go run main.go

build:
	go build -o app ./main.go
