.PHONY:build
build:
	@echo "build program..."
	go build -o bin/txtservice ./cmd/...

.PHONE:clean
clean:
	@rm -r bin

