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

cd %START_DIR%

echo "Running the program"
src\comic-hero.exe -config-dir=..\config -config-file=local-config.json
