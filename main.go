package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := []string{}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// 隠しディレクトリをスキップする
			if info.Name()[0] == '.' {
				return fs.SkipDir
			}
			return nil
		}
		files = append(files, path)
		return nil
	})
	fmt.Println(files)
}
