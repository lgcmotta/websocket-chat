.PHONY: clean build zip

clean:
	rm -rf main*

build:
	GOOS=linux GOARCH=amd64 go build -o main connect.go

zip:
	zip main.zip main 