BINARY_NAME=MarkDown.tar.xz
APP_NAME=MarkDown
VERSION=0.0.1

build:
	fyne package --appVersion 1.0.0 --name MarkDown --icon icon.png --src ./cmd/fyne-md/ --release

run:
	go run ./cmd/fyne-md/main.go

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

test:
	go test -v ./...
