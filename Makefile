server:
	@echo "Starting Go server in dev mode..."
	@APP_ENV="dev" sh -c 'fuser -k 8000/tcp 2>/dev/null || true; go run .'
