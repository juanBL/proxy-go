current-dir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: build
build:
	docker-compose build

.PHONY: start
start:
	docker-compose up -d

.PHONY: sql
sql:
	docker-compose exec mysql sh

##mysql -u zenrows -pzenrows -D zenrows

.PHONY: rm-mysql
rm-mysql:
	docker volume rm zenrows-proxy_mysql_data
