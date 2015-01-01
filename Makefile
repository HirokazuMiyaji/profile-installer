.PHONY: all

all:
	GOOS=darwin GOARCH=amd64 go build -o installer

clean:
	rm installer
