.PHONY: migrate-up migrate-down rest-fiber es

migrate-up:
	migrate -database postgres://goappcore:goappcore@localhost:5432/goappcore?sslmode=disable -path ./examples/rest-fiber/internal/migrations -verbose up

migrate-down:
	migrate -database postgres://goappcore:goappcore@localhost:5432/goappcore?sslmode=disable -path ./examples/rest-fiber/internal/migrations -verbose down

rest-fiber:
	CONFIG_FILE_PATH="$$(pwd)/config" && \
	cd examples/rest-fiber && \
	CONFIG_FILE_PATH=$$CONFIG_FILE_PATH go run main.go

es:
	CONFIG_FILE_PATH="$$(pwd)/config" && \
	cd examples/elasticsearch-implementation && \
	CONFIG_FILE_PATH=$$CONFIG_FILE_PATH go run main.go
