#!/bin/bash

echo "🔨 Building binaries for all major platforms..."

GOOS=linux   GOARCH=amd64 go build -o dist/passy-linux-amd64   main.go
GOOS=darwin  GOARCH=amd64 go build -o dist/passy-macos-intel   main.go
GOOS=darwin  GOARCH=arm64 go build -o dist/passy-macos-arm     main.go
GOOS=windows GOARCH=amd64 go build -o dist/passy-windows.exe   main.go

echo "✅ Done! Binaries saved in ./dist"
