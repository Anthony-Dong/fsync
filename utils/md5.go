package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func MD5File(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()
	result := md5.New()
	if _, err := io.Copy(result, file); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", result.Sum(nil))
}

func MD5(file []byte) string {
	return fmt.Sprintf("%x", md5.Sum(file))
}
