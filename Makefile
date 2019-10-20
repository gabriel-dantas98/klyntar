build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/color-ls
