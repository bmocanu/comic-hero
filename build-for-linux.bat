@echo off
set START_DIR=%cd%

cd src

echo "Removing all exe builds in ./src"
del *.exe
echo "Removing the linux build in ./src"
del comic-hero

echo "Setting build for Linux - AMD64"
set GOOS=linux
set GOARCH=amd64

echo "Building the program"
go build

echo "Building finished"
cd %START_DIR%
