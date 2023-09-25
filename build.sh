#!/bin/bash

set -e

rm -rf bin

function go_build(){
  # CGO_ENABLED=0 GOOS=windows GOARCH=amd64
  # CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
  # CGO_ENABLED=0 GOOS=darwin GOARCH=arm
  # CGO_ENABLED=0 GOOS=linux GOARCH=amd64
  binary="fsync"
  if [ "$1" == "windows" ]; then
    binary="fsync.exe"
  fi
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2  go build -v -ldflags "-s -w" -o "bin/$1_$2/$binary" main.go
}

function format_golang_file () {
  project_dir=$(realpath "$1")
	# shellcheck disable=SC2044
	for elem in $(find "${project_dir}" -name '*.go'); do
		gofmt -w "${elem}"  > /dev/null 2>&1;
		goimports -w -srcdir "${project_dir}" -local "$2" "${elem}" > /dev/null 2>&1;
	done
}

if [ "$1" == "format" ]; then
    format_golang_file . "github.com/anthony-dong/fsync"
    exit 0
fi

if [ "$1" == "install" ];then
    go build -v -o "$(go env GOPATH)/bin/fsync" main.go
    exit 0
fi

go_build windows amd64
go_build darwin amd64
go_build darwin arm64
go_build linux amd64