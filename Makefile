# Auto generate a new model schema with Ent
model:
	cd src; go run -mod=mod entgo.io/ent/cmd/ent new ${model}

# Generate code based on model schemas
generate:
	cd src; go generate ./ent

# Run unit tests
test:
	cd src/tests && go test -v ./...

# Start the docker container and orchestrate containers
start:
	docker compose up -d

# Stop the docker container
stop:
	docker compose down

# Launch ngrok and listen to webhooks
listen:
	ngrok http 8080