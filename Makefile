.PHONY: ci clean test build


ci: clean prep gen test build

build:
	go build github.com/gopheracademy/congo
	cd ui && npm run build && cd ..

build-cli:
	go build github.com/gopheracademy/congo/client/congo-cli

clean:
	rm -rf app/
	rm -rf client/
	rm -rf js/
	rm -rf models/
	rm -rf schema/
	rm -rf swagger/

prep:
	go get -u github.com/goadesign/goa/...
	go get -u github.com/goadesign/gorma/...

gen:
	goagen --design github.com/gopheracademy/congo/design app
	goagen --design github.com/gopheracademy/congo/design js -o ui/app/gen
	goagen --design github.com/gopheracademy/congo/design client
	goagen --design github.com/gopheracademy/congo/design schema
	goagen --design github.com/gopheracademy/congo/design swagger
	goagen --design github.com/gopheracademy/congo/design gen --pkg-path=github.com/bketelsen/gorma



test:
	gb list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' github.com/gopheracademy/congo/... | awk 'NF'  |  xargs gb test

run:
	SESSION_SECRET=ASDF1234ASDFLKAJSDF  TWITTER_KEY=q7ZO1zxazOpAUT06aiTrF83Up  TWITTER_SECRET=uogbiWtHHQQ1nl3OfPOe92vZkt7YtgYgQWlaxGXSj3tqrhnSNC CONGO_ENVIRONMENT=dev CONGO_PORT=8080 ./congo
 
