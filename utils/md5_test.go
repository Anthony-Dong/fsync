package utils

import (
	"testing"
)

func TestMd5File(t *testing.T) {
	t.Log(MD5File("main.go"))
}

func TestMd5(t *testing.T) {
	t.Log(MD5([]byte("main.go")))
}
