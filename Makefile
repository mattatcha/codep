build:
	go build -ldflags "-X main.Version $(shell git describe --tags)" -o codep
