.PHONY: migrate migrate-dry-run \
	generate-queries generate-query-interfaces generate-query-mock \
	watch \
	bun bun-install \
	bundle-js bundle-css \
	bundle-stimulus bundle-turbo bundle-stimulus-use \
	bundle-sakura-css \
	build

.SILENT: bun

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

bun:
	command -v bun > /dev/null || (echo "bun is required to build frontend assets" && exit 1)

bun-install: bun
	bun install

node_modules: bun-install

bundle-stimulus: node_modules
	bun build node_modules/@hotwired/stimulus/dist/stimulus.js \
		--outdir ./asset/js/ \
		--sourcemap=external \
		--minify

bundle-turbo: node_modules
	bun build node_modules/@hotwired/turbo/dist/turbo.es2017-esm.js \
		--outdir ./asset/js/ \
		--sourcemap=external \
		--entry-naming=turbo.[ext] \
		--minify

bundle-stimulus-use: node_modules
	bun build node_modules/stimulus-use/dist/index.js \
		--outdir ./asset/js/ \
		--sourcemap=external \
		--entry-naming=stimulus-use.[ext] \
		--external @hotwired/stimulus \
		--minify

bundle-sakura-css: node_modules
	cp node_modules/sakura.css/css/sakura.css ./asset/css/sakura.css

bundle-js: bundle-stimulus bundle-turbo bundle-stimulus-use
bundle-css: bundle-sakura-css
bundle-assets: bundle-js bundle-css
