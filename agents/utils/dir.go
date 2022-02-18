package utils

import (
	"os"
	"path/filepath"
)

func Mkdir(path string)  {
	os.MkdirAll(path, os.ModePerm)
}

func MkPdir(path string)  {
	dirName := filepath.Dir(path)
	Mkdir(dirName)
}