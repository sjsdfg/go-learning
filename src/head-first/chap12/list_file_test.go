package chap12

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestListingFile(t *testing.T) {
	listFile("D:/DDU")
}

func listFile(path string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range dir {
		if info.IsDir() {
			listFile(filepath.Join(path, info.Name()))
		} else {
			fmt.Println(path + info.Name())
		}
	}
}
