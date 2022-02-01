package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/fs"
	"os"
	"path"
	"time"
)

func main() {
	dir, err := dtDir()
	if err != nil {
		panic(err)
	}
	name := time.Now().Format("2006_01_02")
	f, err := createDirAndFile(dir, name)
	if err != nil {
		panic(err)
	}
	/*
		 fi, err := os.Stat(f)
		 if err != nil {
			 panic(err)
		}
		if fi.Size() == 0 {

		}
	*/
	fmt.Println(f)
}

func dtDir() (string, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := path.Join(dir, ".dt")
	return path, nil
}

func createDirAndFile(dir, name string) (string, error) {
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return "", err
	}
	fPath := path.Join(dir, name)
	f, err := os.OpenFile(fPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return "", err
	}
	return fPath, nil
}
