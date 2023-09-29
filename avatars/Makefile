.PHONY: compose down

help:
	@echo "make compose    -- deploy demo with Compose & Docker"

compose:
	docker compose up --build --wait
	docker compose watch

down:
	docker compose down --timeout=0 --volumes --remove-orphans
endif
