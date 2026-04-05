up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

mod:
	go mod tidy

api-run:
	go run ./cmd/api

cron-run:
	go run ./cmd/scheduler

worker-run:
	go run ./cmd/worker