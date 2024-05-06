# Generate seed data in local development
seed-dev:
	@echo "Running development seeds..."
	docker compose exec app go run scripts/seeds/dev/seed.go

# Generate seed data in production environment
seed-prod:
	@echo "Running production seeds..."
	docker compose exec app go run scripts/seeds/prod/seed.go

# Auto generate a new model schema with Ent
model:
	@echo "Generate a new schema..."
	cd src/api; go run -mod=mod entgo.io/ent/cmd/ent new ${model}

# Generate code based on model schemas
generate:
	@echo "Generate code based on schemas..."
	cd src; go generate ./api/ent

# Run unit tests
test:
	@echo "Running tests..."
	cd src/tests && go test -v ./...

# Start the docker container and orchestrate containers
start:
	@echo "Starting docker container..."
	docker compose up -d

# Stop the docker container
stop:
	@echo "Shutting down docker container..."
	docker compose down