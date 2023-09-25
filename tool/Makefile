lint:
	golangci-lint run

lint-docker:
	docker run --rm -v .:/app -v ~/.cache/golangci-lint/v1.53.3:/root/.cache -w /app golangci/golangci-lint:v1.53.3 golangci-lint run 

lint-docker-build:
	docker build --target lint --output type=cacheonly .
