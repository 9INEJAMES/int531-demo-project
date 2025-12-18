run:
	docker compose up --build -d
stop:
	docker compose down
monitor:
	docker compose -f monitoring.compose.yml up -d