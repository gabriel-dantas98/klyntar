build:
	go mod download
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/color-ls
	chmod +x bin/color-ls
	tar -cvf bin/colorls.tar bin/color-ls

install:
	rm -f /usr/bin/color-ls
	sudo mv bin/color-ls /usr/bin/
