include .env
GO_BIN := $(shell which go)
BUN_BIN := $(shell which bun)
TEMPL_BIN := $(shell which templ)

PROJECT_NAME := $(shell basename $(CURDIR))
STATIC_DIR := static
BUILD_DIR := build
TEMP_DIR := tmp

TAILWIND_CONFIG := 'module.exports = {\n\
  content: ["./views/**/*.templ"],\n\
  theme: {\n\
    extend: {}\n\
  },\n\
  plugins: []\n\
}'

ALPINE_URL := https://cdn.jsdelivr.net/npm/alpinejs@latest/dist/cdn.min.js

check-deps:
	@test -n "$(GO_BIN)" || (printf "✗ Go not installed\n" && exit 1)
	@test -n "$(BUN_BIN)" || (printf "✗ Bun not installed\n" && exit 1)
	@test -n "$(TEMPL_BIN)" || (printf "✗ templ not installed\n" && exit 1)

create-dirs:
	@mkdir cmd
	@mkdir -p $(STATIC_DIR)/js
	@mkdir -p $(STATIC_DIR)/css
	@mkdir -p views/{components,layouts,pages}
	@mkdir -p {handlers,middlewares,models,services,utils}
	@mkdir -p $(BUILD_DIR)

setup-go:
	@$(GO_BIN) mod init $(PROJECT_NAME) 2>/dev/null || true
	@$(GO_BIN) mod tidy >/dev/null 2>&1

setup-tailwind:
	@$(BUN_BIN) install tailwindcss@latest >/dev/null 2>&1
	@echo -e $(TAILWIND_CONFIG) > tailwind.config.js
	@echo -e '@import "tailwindcss"' > $(STATIC_DIR)/css/input.css

download-alpine:
	@curl -s $(ALPINE_URL) -o $(STATIC_DIR)/js/alpine.min.js

templ:
	@$(TEMPL_BIN) generate

css:
	@$(BUN_BIN) exec "tailwindcss -i $(STATIC_DIR)/css/input.css -o $(STATIC_DIR)/css/styles.css --minify"

serve: templ css
	@$(GO_BIN) run ./cmd/main.go serve --http ${ROUTE}

clean:
	@rm -rf $(BUILD_DIR)
	@rm -rf $(TEMP_DIR)
	@find . -type f -name "*_templ.go" -delete

watch:
	@find views -type f -name "*.templ" | entr -r make serve

init: check-deps create-dirs setup-go setup-tailwind download-alpine

build: check-deps
	@$(GO_BIN) build -o $(BUILD_DIR)/$(PROJECT_NAME) ./main.go

