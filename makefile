BINARY_DIR ::= bin
BINARY_NAME ::= embed-exp
CMD_DIR ::= cmd
BUILD_FLAGS ::= -v -x --buildvcs

.PHONY: all build-setup tidy build clean install

all: setup-build tidy build
	@echo "Running all targets"

build-setup: mkdir-bin
	@echo "Setting up the project..."
	mkdir -p $(BINARY_DIR)

tidy:
	@echo "Tidying up dependencies..."
	go mod tidy -e -x -v

build: build-setup
	@echo "Building program..."
	go build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go

install:
	@echo "Installing program..."
	#go install $(BUILD_FLAGS) -o $(BINARY_NAME) $(CMD_DIR)/main.go

clean:
	@echo "Cleaning build..."
	rm -f $(BINARY_DIR)/$(BINARY_NAME)
	rm -fd $(BINARY_DIR)
