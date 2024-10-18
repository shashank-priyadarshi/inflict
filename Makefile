all: windows darwin linux

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/inflict-windows-amd64.exe .
	GOOS=windows GOARCH=386 go build -o bin/inflict-windows-386.exe .
darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/inflict-darwin-amd64 .
linux:
	GOOS=linux GOARCH=amd64 go build -o bin/inflict-linux-amd64 .
	GOOS=linux GOARCH=386 go build -o bin/inflict-linux-386 .
