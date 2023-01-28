pmu:
	npx prisma migrate dev -n init --schema ./db/schema.prisma --skip-generate --create-only
	find ./db/migrations -type f -exec sed -i 's#"##g' {} +
	rm ./db/migrations/*.toml
	rm ./db/migrations/migration.sql
	cp ./db/migrations/**/*.sql ./db/migrations
	rm -rf node_modules
	rm -rf package.json
	rm -rf package-lock.json
	rm -r ./db/migrations/*_init

sqlc:
	rm -r ./db/sqlc
	sqlc generate
migratecreate:
	migrate create -ext sql -dir db/migrations -seq init
