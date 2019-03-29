package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

var (
	s = "│"
	m = "├───"
	e = "└───"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(output io.Writer, currDir string, printFiles bool) error {
	recursionPrintService("", output, currDir, printFiles)
	return nil
}

func recursionPrintService(prependingString string, output io.Writer, currDir string, printFiles bool) {
	fileObj, err := os.Open(currDir)
	defer fileObj.Close()
	if err != nil {
		panic("Couild not open file")
	}
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	if err != nil {
		panic("Couild not read current dirrectory")
	}
	l := len(files)
	for _, file := range files {
		i := 1
		if !file.IsDir() && printFiles {
			if file.Size() > 0 {
				if l > i+1 {
					fmt.Fprintf(output, prependingString+m+"%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, prependingString+e+"%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if l > i+1 {
					fmt.Fprintf(output, prependingString+m+"%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, prependingString+e+"%s (empty)\n", file.Name())
				}
			}
		} else {
			newDir := path.Join(currDir, file.Name())
			fmt.Fprintf(output, prependingString+e+"%s\n", file.Name())
			if l > i+1 {
				prependingString += "\t" + s
			} else {
				prependingString += "\t"
			}
			recursionPrintService(prependingString, output, newDir, printFiles)
		}
	}
}
