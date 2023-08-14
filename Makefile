up: cache
	docker-compose up -d --build

down:
	docker compose down

schemas:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./internal/platform/repository/sql/ent ./internal/platform/repository/sql/schema

test:
	go test -v ./...

cache:
	go clean --modcache
	go mod tidy
	go mod vendor
	modvendor -copy='**/*.c **/*.h **/*.a'

mocks:
	mockgen -source=internal/domain/repository/credential_repository.go -destination=internal/domain/repository/mocks/mock_credential_repository.go