package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	os.Exit(recurse(dir))
}

func recurse(dir string) int {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("error: %s", err)
		return 1
	}
	var errint int
	for _, fs := range fs {
		if strings.HasPrefix(fs.Name(), ".") {
			continue
		}
		fp := filepath.Join(dir, fs.Name())
		stat, err := os.Stat(fp)
		if err != nil {
			fmt.Printf("error with %s : %s", fp, err)
			errint = 1
			continue
		}
		fmt.Printf("%s\t%s\n", stat.Mode().String(), fp)
		if stat.IsDir() {
			if ret := recurse(fp); ret != 0 {
				errint = ret
			}
		}
	}
	return errint
}
