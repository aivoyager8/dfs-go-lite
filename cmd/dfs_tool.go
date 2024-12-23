
package main

import (
	"fmt"
	"os"
	"github.com/yourusername/dfs-go-lite/dfs"
)

func listFilesAndDirectories(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
}

func main() {
	root := "." // Change this to the directory you want to list
	err := listFilesAndDirectories(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}