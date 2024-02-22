.DEFAULT_GOAL := build

.PHONY: build tidy test run clean

build:
	go build .

tidy:
	go mod tidy

test:
	find . -type d -maxdepth 1 -not -path "./test" -exec go test {} \;

integration-test:
	if [ -d "test" ]; then go test -v ./test; else echo "no integration tests"; fi

clean:
	rm -rf where-is-love
