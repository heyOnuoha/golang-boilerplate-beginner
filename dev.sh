#!/bin/bash

echo "Generating Swagger documentation..."
# Use the full path to swag in the Go bin directory
SWAG_PATH="$HOME/go/bin/swag"

# Check if swag is installed
if [ ! -f "$SWAG_PATH" ]; then
    echo "swag is not installed. Installing..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate Swagger docs
"$SWAG_PATH" init -g main.go -o docs

echo "Swagger documentation updated successfully!"

# Check if air is installed
AIR_PATH="$HOME/go/bin/air"
if [ ! -f "$AIR_PATH" ]; then
    echo "air is not installed. Installing..."
    go install github.com/air-verse/air@latest
fi

# Check if port 8080 is available, if not use 8081
DEFAULT_PORT=8080

if lsof -i:$DEFAULT_PORT > /dev/null 2>&1; then
    echo "Port $DEFAULT_PORT is already in use, using port 8081 instead"
    export PORT=8081
else
    echo "Using default port $DEFAULT_PORT"
    export PORT=$DEFAULT_PORT
fi

echo "Starting development server on port $PORT..."
# Run Air with the configuration file
"$AIR_PATH" -c .air.toml
