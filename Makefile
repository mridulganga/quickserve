build:
	go build -v -o quickserve main.go
	GOOS=windows GOARCH=amd64 go build -v -o quickserve.exe main.go

vendor:
	go mod tidy