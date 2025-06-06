.PHONY: all build test run clean

# Variables
BINARY_NAME=gomon
GO=go
MAIN_PATH=cmd/gomon/main.go

# Version info
VERSION?=0.1.0
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Couleurs pour le formatage
BOLD := $(shell tput bold)
BLUE := $(shell tput setaf 4)
GREEN := $(shell tput setaf 2)
YELLOW := $(shell tput setaf 3)
RED := $(shell tput setaf 1)
RESET := $(shell tput sgr0)

all: test build

# Construction du binaire
build:
	@echo "Building Gomon..."
	$(GO) build -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Exécution des tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Lancement en mode développement
run:
	@echo "Starting Gomon..."
	$(GO) run $(MAIN_PATH)

# Nettoyage
clean:
	@echo "Cleaning..."
	rm -rf bin/
	$(GO) clean

# Installation des dépendances de développement
dev-deps:
	$(GO) install github.com/golang/mock/mockgen@latest
	$(GO) install golang.org/x/lint/golint@latest

# Linting
lint:
	golint ./...
	$(GO) vet ./...

# Pour tester manuellement les endpoints avec un affichage amélioré
test-endpoints:
	@echo "Testing Gomon endpoints..."
	@bash scripts/test-endpoints.sh || echo "Error: Server must be running. Start with 'make run' first"

# Pour lancer le serveur
run-server:
	go run cmd/gomon/main.go

