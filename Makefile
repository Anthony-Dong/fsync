.PHONY: all install format

all: install

build: format
	@bash build.sh

install: format
	@bash build.sh install

clean:
	rm -rf output
	rm -rf mac_tools.tgz

format:
	@bash build.sh format

test:
	@mkdir -p output
	@go test -v ./...
	@fsync copy --from . --to output --exclude-from .gitignore --exclude .git --log-level debug

tar: clean build
	@tar --exclude=.git --exclude=*.zip --exclude=*.tgz -zcvf fsync.tgz ./*
	@zip -x .git -x *.tgz -x *.zip -r fsync.zip ./*

upload: tar
	gtool upload --decode url --file ./fsync.zip --type software