
test:
	go test -v ./...

app:
	go run .

image:
	docker build --no-cache -t my-golang-app:latest .

fmt:
	go fmt ./...

