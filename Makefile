# Auto generate a new model schema with Ent
model:
	cd src; go run -mod=mod entgo.io/ent/cmd/ent new ${model}

# Generate code based on model schemas
generate:
	cd src; go generate ./ent

# Start the docker container and orchestrate containers
start:
	docker compose up -d

# Stop the docker container
stop:
	docker compose down