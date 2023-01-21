build-go:
	go mod init github.com/nikhilakki/toml-in-pyrsgo
	go get
	go mod tidy

go:
	go run main.go

build-py:
	pip install toml

py:
	python3 main.py	