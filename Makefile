all:

dbuild:
	docker build -t tiktok-go .

drun:
	docker run -dp 8080:8080 tiktok-go

dup:
	docker compose up -d

ddown:
	docker compose down
