@echo off
set START_DIR=%cd%

cd src

echo "Removing all exe builds in ./src"
del *.exe

echo "Setting build for Windows - AMD64"
set GOOS=windows
set GOARCH=amd64

echo "Building the program"
go build

echo "Setting up the env for: http://localhost:8080/comic-hero"
set LISTEN_ADDRESS=localhost
set LISTEN_PORT=8080
set CONTEXT_PATH=/comic-hero

echo "Running the program"
comic-hero.exe

cd %START_DIR%
