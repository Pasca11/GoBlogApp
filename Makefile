.SILENT:

export GOPROXY=https://proxy.golang.org

export GOSUMDB=on

.PHONY:

LOCAL_BIN=$(CURDIR)/bin
POSTGRES_DB_CONTAINER_NAME="blog-postgres"
MONGODB_CONTAINER_NAME="blog-mongodb"
DATANODE_CONTAINER_NAME="blog-datanode"
GRAYLOG_CONTAINER_NAME="blog-graylog"
ELASTICSEARCH_CONTAINER_NAME="blog-elasticsearch"

POSTGRES_HOST=localhost
POSTGRES_USER=amir
POSTGRES_DB=blog
POSTGRES_SLL_MODE=disable
POSTGRES_PORT=6543

.PHONY: prepare-local-env
prepare-local-env:
	docker rm --force "$(POSTGRES_DB_CONTAINER_NAME)" || true
	docker rm --force "$(MONGODB_CONTAINER_NAME)" || true
	docker rm --force "$(DATANODE_CONTAINER_NAME)" || true
	docker rm --force "$(GRAYLOG_CONTAINER_NAME)" || true
	docker rm --force "$(ELASTICSEARCH_CONTAINER_NAME)" || true
	docker-compose --env-file .env -f ./.deploy/build/docker-compose.yaml up -d
	sleep 5

.PHONY:
run:
	go run cmd/blog/main.go