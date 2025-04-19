#!/bin/bash

set -e

PROJECT_NAME="go-init"
MODULE_NAME="github.com/skriptvalley/${PROJECT_NAME}"

echo "🚀 Setting up Go project: $PROJECT_NAME"

# Create folder structure
mkdir -p cmd/server
mkdir -p internal/server
mkdir -p pkg
mkdir -p api
mkdir -p configs
mkdir -p deployments
mkdir -p scripts

# Create go.mod
cd $(dirname "$0")
go mod init $MODULE_NAME

# Create .gitignore
cat <<EOL > .gitignore
# Binaries
/bin/
/build/
*.exe
*.out

# IDEs
.vscode/
.idea/

# Env files
.env

# OS
.DS_Store
EOL

# Create README.md
cat <<EOL > README.md
# ${PROJECT_NAME}

🚀 A production-ready Golang microservice starter kit.

## Structure
- \`cmd/\`: Entrypoint binaries
- \`internal/\`: Internal app logic
- \`pkg/\`: Shared packages
- \`configs/\`: Config and environment files
- \`deployments/\`: Docker & K8s configs

## Quick Start

\`\`\`bash
make run
\`\`\`
EOL

# Create Makefile
cat <<EOL > Makefile
run:
	go run ./cmd/server/main.go

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .

EOL

# Create .env file
cat <<EOL > .env
PORT=8080
LOG_LEVEL=debug
EOL

# Create cmd/server/main.go
cat <<'EOL' > cmd/server/main.go
package main

import (
	"log"

	"github.com/skriptvalley/go-init/internal/server"
)

func main() {
	log.Println("Starting server...")
	server.Start()
}
EOL

# Create internal/server/server.go with Zap logger
cat <<'EOL' > internal/server/server.go
package server

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting HTTP server", "port", port)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		sugar.Infow("Health check endpoint hit")
		fmt.Fprintln(w, "ok")
	})

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		sugar.Fatalw("Server failed", "error", err)
	}
}
EOL

echo "✅ Go project initialized."
