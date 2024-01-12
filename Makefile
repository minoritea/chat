.PHONY: migrate migrate-dry-run \
	generate-queries generate-query-interfaces generate-query-mock \
	generate-session-mock \
	watch \
	build

migrate:
	go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.16.13 ./chat.db < ./database/schema.sql

migrate-dry-run:
	go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.16.13 --dry-run ./chat.db < ./database/schema.sql

generate-queries:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.24.0 generate
	$(MAKE) generate-query-interfaces
	$(MAKE) generate-query-mock

generate-query-interfaces:
	go run github.com/vburenin/ifacemaker@v1.2.1 \
		-f database/queries.sql.go \
		-s Queries \
		-i Querier \
		-p database \
		-o database/interface.go

generate-query-mock:
	go run go.uber.org/mock/mockgen@v0.4.0 \
		-source database/interface.go \
		-destination test/mock/database/querier.go \
		-package database

generate-session-mock:
	go run go.uber.org/mock/mockgen@v0.4.0 \
		-destination test/mock/session/store.go \
		-package session \
		github.com/gorilla/sessions Store

watch:
	go run github.com/makiuchi-d/arelo@latest \
		-p '**/*.go' \
		-p 'template/*.tmpl' \
		-p 'asset/js/*.js' \
		-p 'asset/css/*.css' \
		-i '**/_test.go' \
		-i '**/.*' \
		-- go run main.go

build:
	go build -ldflags "-s -w -X 'main.version=$(shell git rev-parse --short HEAD)'" -trimpath
