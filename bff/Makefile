pmu:
	npx prisma migrate dev -n init --schema ./db/schema.prisma --skip-generate --create-only
	find ./db/migrations -type f -exec sed -i 's#"##g' {} +
	rm ./db/migrations/*.toml
	rm -rf node_modules
	rm package.json
	rm package-lock.json
sqlc:
	sqlc generate
migratecreate:
	migrate create -ext sql -dir db/migrations -seq init
