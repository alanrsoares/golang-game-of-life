# Makefile for building and running a Go project.

# Set the output binary name.
BINARY_NAME=myapp

# Set the source directory.
SRC_DIR=./src

# Default make command should build the project.
all: build

# Build the project.
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(SRC_DIR)

# Run the project.
run:
	@echo "Running..."
	go run $(SRC_DIR)

# Run the tests.
test:
	@echo "Testing..."
	go test -v $(SRC_DIR)

# Clean up the binary.
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)

# "make build" will compile the app and produce a binary.
# "make run" will run the app using 'go run'.
# "make clean" will remove the binary and clean up the project.
