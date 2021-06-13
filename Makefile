setup: ## セットアップ
	cp config.example.yml config.yml
	cp config.testing.example.yml config.testing.yml
	npm run prod

build: ## コンテナビルド
	docker-compose build --no-cache mysql pma redis mailhog

start: ## コンテナ起動
	docker-compose up -d mysql pma redis mailhog

start-swagger: ## Swaggerコンテナ起動
	docker-compose up -d swagger-editor swagger-ui swagger-api

stop: ## コンテナ停止
	docker-compose stop

testing: ## テスト
	rm -f fiber.testing.sqlite
	go test -v ./test
