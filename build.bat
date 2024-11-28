cd /d %~dp0
@REM set CC=arm-none-linux-gnueabihf-gcc
@REM set CXX=arm-none-linux-gnueabihf-g++
@REM set CGO_ENABLED=1
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -o ./bin/lumenim ./cmd/lumenim
