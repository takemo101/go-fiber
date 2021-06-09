setup: ## セットアップ
	cp config.example.yml config.yml
	npm run prod

build: ## コンテナビルド
	docker-compose build --no-cache mysql pma redis mailhog

start: ## コンテナ起動
	docker-compose up -d mysql pma redis mailhog

start-swagger: ## Swaggerコンテナ起動
	docker-compose up -d swagger-editor swagger-ui swagger-api

stop: ## コンテナ停止
	docker-compose stop
