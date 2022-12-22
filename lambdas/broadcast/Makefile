.PHONY: clean build

clean:
	rm -rf main*

build:
	GOOS=linux GOARCH=amd64 go build -o main broadcast.go && zip main.zip main