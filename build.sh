#!/bin/bash

PROJECT_ROOT=$(cd "$(dirname "$0")" && pwd)
echo "Project root directory: $PROJECT_ROOT"
cd "$PROJECT_ROOT" || exit

if [ ! -f "go.mod" ]; then
    echo "Initializing root Go module..."
    go mod init urlshortner
else
    echo "Root Go module already initialized."
fi

# Tidy and test all modules
find . -name 'go.mod' -execdir go mod tidy \;
find . -name 'go.mod' -execdir go test ./... \;

# Build the project
echo "Building the project..."
go build -o bin/url_shortner_main ./src/cmd/url_shortner_main/
# Need this since I am coding on mac and using linux docker
GOOS=linux GOARCH=amd64 go build -o bin/url_shortner_main_for_docker ./src/cmd/url_shortner_main/

# Check if the build was successful
if [ -f "bin/url_shortner_main" ]; then
    echo "Build successful. Binary created at $PROJECT_ROOT/bin/url_shortner_main"
else
    echo "Build failed."
fi

