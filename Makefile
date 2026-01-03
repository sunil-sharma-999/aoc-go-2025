SRC_DIR := ./src
BUILD_DIR := ./dist

# Find all main.go files in src/day* directories
MAIN_FILES := $(shell find $(SRC_DIR) -name main.go)

.PHONY: all
all: $(MAIN_FILES:$(SRC_DIR)/%/main.go=$(BUILD_DIR)/%)

# Build rule
$(BUILD_DIR)/%: $(SRC_DIR)/%/main.go
	@echo "Building $<"
	mkdir -p $(dir $@)
	go build -o $@ $<

.PHONY: run
run: $(BUILD_DIR)/day$(word 2, $(MAKECMDGOALS))
	@ \
	day=$(word 2, $(MAKECMDGOALS)); \
	args=$(filter-out $@ $(day),$(MAKECMDGOALS)); \
	$(BUILD_DIR)/day$$day $$args

# Prevent make from treating day number as a target
%:
	@:
