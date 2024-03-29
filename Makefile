test:
	go build -v -o dist/go-cart main.go
#  ./scripts/generate_test_pki.bundle.sh

test-run:
	go build -v -o dist/go-cart main.go
	./dist/go-cart -config ./config/config.yml.example

test-run-local:
	go build -v -o dist/go-cart main.go
	./dist/go-cart -config ./config.yml

build:
	go build -v -o dist/go-cart main.go

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o dist/go-cart-linux-amd64 main.go

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -v -o dist/go-cart-darwin-amd64 main.go

run:
	go run main.go