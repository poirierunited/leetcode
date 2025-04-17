.PHONY: run run-all test test-all clean

# Default problem to run if none specified
PROBLEM ?= 

# Run a specific problem
run:
	@if [ -z "$(PROBLEM)" ]; then \
		echo "Please specify a problem: make run PROBLEM=problemname"; \
		exit 1; \
	fi
	@echo "Running $(PROBLEM)..."
	@go run cmd/problems/$(PROBLEM)/main.go

# Run all problems
run-all:
	@for dir in cmd/problems/*/; do \
		if [ -d "$$dir" ]; then \
			problem=$$(basename $$dir); \
			echo "Running $$problem..."; \
			go run $$dir/main.go; \
			echo "-------------------"; \
		fi \
	done

# Test a specific problem
test:
	@if [ -z "$(PROBLEM)" ]; then \
		echo "Please specify a problem: make test PROBLEM=problemname"; \
		exit 1; \
	fi
	@echo "Testing $(PROBLEM)..."
	@go test ./cmd/problems/$(PROBLEM)/...

# Test all problems
test-all:
	@for dir in cmd/problems/*/; do \
		if [ -d "$$dir" ]; then \
			problem=$$(basename $$dir); \
			echo "Testing $$problem..."; \
			go test ./$$dir/...; \
			echo "-------------------"; \
		fi \
	done

# Clean build artifacts
clean:
	@go clean
	@find . -name "*.test" -delete 