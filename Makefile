up: cache
	docker-compose up -d --build

schemas:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./internal/platform/repository/sql/ent ./internal/platform/repository/sql/schema

test:
	go test -v ./...

cache:
	go mod tidy
	go mod vendor