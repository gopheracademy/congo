.PHONY: ci clean test build


ci: clean prep gen test build

build:
	go build github.com/gopheracademy/congo


clean: pgdown 
	rm -rf app/
	rm -rf client/
	rm -rf js/
	rm -rf schema/
	rm -rf swagger/


prep:
	go get -u github.com/goadesign/goa/...
	go get -u github.com/goadesign/gorma/...

gen:
	goagen --design github.com/gopheracademy/congo/design app
	goagen --design github.com/gopheracademy/congo/design js 
	goagen --design github.com/gopheracademy/congo/design client
	goagen --design github.com/gopheracademy/congo/design schema
	goagen --design github.com/gopheracademy/congo/design swagger
	goagen --design github.com/gopheracademy/congo/design gen --pkg-path=github.com/goadesign/gorma

test:
	go test ./...

pgup:
	docker-compose -f dc-postgres.yml up -d

pgdown:
	docker-compose -f dc-postgres.yml down

pgclean:
	docker-compose -f dc-postgres.yml rm

appdev:
	CONGO_DEBUG=true CONGO_DB_HOST=docker.local CONGO_DB_USERNAME=congo CONGO_DB_NAME=congo CONGO_DB_PORT=5432 CONGO_DB_PASSWORD=congopass ./congo

run: pgup build appdev
 
key_checks: 
	which ssh-keygen > /dev/null
	which openssl > /dev/null

keys: key_checks
	mkdir keys
	ssh-keygen -t rsa -b 4096 -f ./keys/congo.rsa -t rsa -N ''
	openssl genrsa -out ./keys/congo.key 4096
	openssl req -new -key ./keys/congo.key -out ./keys/congo.csr
