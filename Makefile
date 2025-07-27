start:
	@echo "Starting product service..."
	docker compose up --build -d

stop:
	@echo "Stopping product service..."
	docker compose down