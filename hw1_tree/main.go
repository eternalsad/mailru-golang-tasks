package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	path := "testdata"
	printFiles := false
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(file io.Writer, path string, printFiles bool) error {
	prefix := "" // no prefix on first iteration
	// fmt.Println(printFiles)
	return recursion(file, path, printFiles, prefix)
}

const (
	pr     = "├───"
	lastPR = "└───"
	add    = "    "
)

func recursion(out io.Writer, path string, printFiles bool, indent string) error {
	dirEntry, err := ioutil.ReadDir(path)
	// fmt.Println(path)
	if err != nil {
		return err
	}
	// pd := "\t|"
	// if dirEntry[0].IsDir() {
	// 	pd = "|"
	// }
	// indent += "\t"
	for i := 0; i < len(dirEntry)-1; i++ {
		entry := dirEntry[i]
		prefix := "├───"
		if printFiles && !entry.IsDir() {
			if entry.Size() == 0 {
				fmt.Fprintln(out, indent+prefix+entry.Name()+" (empty)")
			} else {
				fmt.Fprintln(out, indent+prefix+entry.Name()+fmt.Sprintf(" (%vb)", entry.Size()))
			}
		}
		if entry.IsDir() {
			if path == "" {
				indent = ""
			}
			fmt.Fprintln(out, indent+prefix+entry.Name())
			recursion(out, path+"/"+entry.Name(), printFiles, indent+"│\t")
		}
	}
	last := dirEntry[len(dirEntry)-1]
	prefix := "└───"
	if printFiles && !last.IsDir() {
		if last.Size() == 0 {
			fmt.Fprintln(out, indent+prefix+last.Name()+" (empty)")
		} else {
			fmt.Fprintln(out, indent+prefix+last.Name()+fmt.Sprintf(" (%vb)", last.Size()))
		}
	}
	if last.IsDir() {
		fmt.Fprintln(out, indent+prefix+last.Name())
		// indent += "\t"
		recursion(out, path+"/"+last.Name(), printFiles, indent+"\t")
	}
	return nil
}

// просто префикс увеличивать на один таб а вот уже печатать палку или нет решать внутри ифа
