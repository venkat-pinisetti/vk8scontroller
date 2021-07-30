info:
	echo "Executing Sample Controller code....."

build:
	go build -o vcontroller ./cmd/main.go

run:
	go run ./cmd/main.go

all: info build
