BINDIR := /usr/bin

main:
	go build ./cmd/main.go

install:
	install -Dm755 main "$(DESTDIR)/$(BINDIR)/z85m"
	strip "$(DESTDIR)/$(BINDIR)/z85m"

test:
	go test ./...
