package main

import (
	"fmt"
	"github.com/yusenfei/memcache/inmemcache/inmemcache"
	"io/fs"
)

func main() {
	rootFS := inmemcache.New()

	err := rootFS.MkdirAll("dir1/dir2", 0777)
	if err != nil {
		panic(err)
	}

	err = rootFS.WriteFile("dir1/dir2/f1.txt", []byte("incinerating-unsubstantial"), 0755)
	if err != nil {
		panic(err)
	}

	err = fs.WalkDir(rootFS, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	content, err := fs.ReadFile(rootFS, "dir1/dir2/f1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)
}
