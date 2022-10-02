main:
	go build ./cmd/main.go

install:
	install -Dm755 main "$(DESTDIR)/usr/bin/z85m"
