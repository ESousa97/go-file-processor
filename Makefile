.PHONY: build test bench clean generate-data

BINARY_NAME=fileproc
DATA_DIR=test_data
LARGE_CSV=$(DATA_DIR)/large_test.csv
LARGE_JSON=$(DATA_DIR)/large_test.json

build:
	go build -o $(BINARY_NAME) ./cmd/main.go

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./internal/processor/...

generate-data:
	mkdir -p $(DATA_DIR)
	go run scripts/gen_data.go $(LARGE_CSV) 100000

clean:
	rm -f $(BINARY_NAME)
	rm -rf $(DATA_DIR)
	go clean
