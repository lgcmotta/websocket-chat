.PHONY: clean build

clean:
	rm -rf main*

build:
	GOOS=linux GOARCH=amd64 go build -o main members.go && zip main.zip main