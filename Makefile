build:
	go build -ldflags "-X main.Version $(git describe --tags)" -o /bin/registrator
