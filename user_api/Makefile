GOOS	 ?= linux
GOARCH  = amd64
ODIR    = deploy/_output

TIMESTAMP = $(shell date +'%Y%m%d%H%M%S')
PWD			 	= $(shell pwd)

IMAGE_NAME = rufuskopral/efishery_user_api
VERSION    = $(shell git show -q --format=%h)

export GO111MODULE ?= on

all: compile test

test:
	go test -v ./...

coverage:
	go tool cover -html=cover.out -o cover.html

bin:
	go build -o $(ODIR)/user_api app/main.go

run: bin
	./$(ODIR)/user_api

mod:
	go mod tidy

compile:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(ODIR)/user_api_docker app/main.go

migration:
	@echo "class Migration < ActiveRecord::Migration[5.1]\n  def up\n  end\n\n  def down\n  end\nend" > db/migrate/$(TIMESTAMP)_migration.rb

db-setup:
	docker run --rm -it --network host -v "$(PWD)/db":/app/db -v "$(PWD)/.env":/app/.env rufuskopral/standalone-migration:1.0.0 db:create
	docker run --rm -it --network host -v "$(PWD)/db":/app/db -v "$(PWD)/.env":/app/.env rufuskopral/standalone-migration:1.0.0 db:migrate

db-migrate:
	docker run --rm -it --network host -v "$(PWD)/db":/app/db -v "$(PWD)/.env":/app/.env rufuskopral/standalone-migration:1.0.0 db:migrate

db-migrate-rollback:
	docker run --rm -it --network host -v "$(PWD)/db":/app/db -v "$(PWD)/.env":/app/.env rufuskopral/standalone-migration:1.0.0 db:rollback

db-seed:
	docker run --rm -it --network host -v "$(PWD)/db":/app/db -v "$(PWD)/.env":/app/.env rufuskopral/standalone-migration:1.0.0 db:seed

docker-build:compile
	docker build -t $(IMAGE_NAME) -f ./deploy/Dockerfile .

docker-push:docker-build
		docker push $(IMAGE_NAME);