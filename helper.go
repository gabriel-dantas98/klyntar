package main

import (
	"os"
)

func checkPath(Path string) bool {
	if _, err := os.Stat(Path); !os.IsNotExist(err) {
		return true
	}
	return false
}
