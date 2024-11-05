-include .env

run:
	go run cmd/main.go

migrate_up:
	migrate -path migrations -database ${PSQL_URI} -verbose up


migrate_down:
	migrate -path migrations -database ${PSQL_URI} -verbose down 1

migrate_force:
	migrate -path migrations -database ${PSQL_URI} force ${VERSION}

migrate_file:
	migrate create -ext sql -dir ./migrations -seq ${FILE_NAME}

sqlc:
	sqlc generate
