build: main.go
	@echo "Building gomovies..."
	@rm -rf target
	@mkdir -p target
	@go build -o target/gomovies main.go

run: build
	@echo "Running gomovies..."
	@./target/gomovies