migrate:
	go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.16.13 ./chat.db < ./schema.sql

migrate-dry-run:
	go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.16.13 --dry-run ./chat.db < ./schema.sql

generate-queries:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.24.0 generate

watch:
	go run github.com/makiuchi-d/arelo@latest \
		-p '**/*.go' \
		-p '**/*.tmpl' \
		-i '**/_test.go' \
		-i '**/.*' \
		-- go run main.go
