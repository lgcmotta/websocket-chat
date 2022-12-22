.PHONY: clean build zip

clean:
	rm -rf main*

build:
	GOOS=linux GOARCH=amd64 go build -o main publish.go

zip:
	zip main.zip main 