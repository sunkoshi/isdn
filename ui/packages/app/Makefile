pgen:
	npx prisma generate --schema ./prisma/schema.prisma
pmu:
	npx prisma migrate dev -n init --schema ./prisma/schema.prisma
test:
	yarn test

.PHONY: pgen pmu test