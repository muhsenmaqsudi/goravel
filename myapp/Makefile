BINARY_NAME=goravelApp

build:
	@go mod vendor
	@echo "Building Goravel..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Goravel built!"

run: build
	@echo "Starting Goravel..."
	@./tmp/${BINARY_NAME} &
	@echo "Goravel started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Goravel..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Goravel!"

restart: stop start