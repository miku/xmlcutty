SHELL := /bin/bash
TARGETS = xmlcutty

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go test -v ./...

imports:
	goimports -w .

fmt:
	go fmt ./...

all: fmt test
	go build

install:
	go install

clean:
	go clean
	rm -f coverage.out
	rm -f $(TARGETS)
	rm -f xmlcutty-*.x86_64.rpm
	rm -f packaging/debian/xmlcutty_*.deb
	rm -f xmlcutty_*.deb
	rm -rf packaging/debian/xmlcutty/usr

cover:
	go get -d && go test -v	-coverprofile=coverage.out
	go tool cover -html=coverage.out

xmlcutty: cmd/xmlcutty/main.go
	go build -o xmlcutty cmd/xmlcutty/main.go

# ==== packaging

deb: $(TARGETS)
	mkdir -p packaging/debian/xmlcutty/usr/sbin
	cp $(TARGETS) packaging/debian/xmlcutty/usr/sbin
	cd packaging/debian && fakeroot dpkg-deb --build xmlcutty .
	mv packaging/debian/xmlcutty*deb .

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/rpm/xmlcutty.spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	./packaging/rpm/buildrpm.sh xmlcutty
	cp $(HOME)/rpmbuild/RPMS/x86_64/xmlcutty*.rpm .
