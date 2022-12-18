.PHONY: run dev build

dev:
	@echo "Running..."
	go run -tags="feature playingcard tradingcard" .
	@echo "Done."

build:
	@echo "Building..."
	go build -o ./bin/app -tags="feature playingcard tradingcard" .
	@echo "Done."

run: build
	@echo "Running..."
	./bin/app
	@echo "Done."

clean:
	@echo "Cleaning..."
	rm -rf ./bin
	@echo "Done."