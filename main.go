package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/alloy/syntax/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("validate-grafana-alloy requires exactly one argument: the directory to validate")
		os.Exit(1)
	}

	alloyFiles, err := getAlloyFiles(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range alloyFiles {
		fileContents, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = parser.ParseFile(file, fileContents)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func getAlloyFiles(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".alloy" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
