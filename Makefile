.PHONY: app-up
app-up:
	docker compose up -d

.PHONY: app-down
app-down:
	docker compose down

.PHONY: app-rebuild
app-rebuild:
	docker compose up -d --build