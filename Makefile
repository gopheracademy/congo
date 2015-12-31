.PHONY: ci clean test build


ci: clean prep gen test build

build: 
	go build github.com/gopheracademy/congo
	go build github.com/gopheracademy/congo/client/congo-cli

clean:
	rm -rf app/
	rm -rf client/
	rm -rf js/
	rm -rf models/*_genmodel.go 
	rm -rf schema/
	rm -rf swagger/

prep:
	go get -u github.com/raphael/goa/...
	go get -u github.com/bketelsen/gorma/...

gen:
	goagen --design github.com/gopheracademy/congo/design app
	goagen --design github.com/gopheracademy/congo/design js
	goagen --design github.com/gopheracademy/congo/design client
	goagen --design github.com/gopheracademy/congo/design schema
	goagen --design github.com/gopheracademy/congo/design swagger
	goagen --design github.com/gopheracademy/congo/design gen --pkg-path=github.com/bketelsen/gorma


# Test all XOR specific packages, skip packages that don't have tests for performance reasons
test:
	gb list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' github.com/gopheracademy/congo/... | awk 'NF'  |  xargs gb test

keys: ; @which ssh-keygen > /dev/null
	mkdir keys
	ssh-keygen -t rsa -b 4096 -f ./keys/congo.rsa -t rsa -N ''
	openssl genrsa -out ./keys/congo.key 4096
	openssl req -new -key ./keys/congo.key -out ./keys/congo.csr