package gobyexample

import (
	"fmt"
	"os"
	"path/filepath"
)

func deferCreateFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func deferWriteFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func deferCloseFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

func ShowDefer() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := deferCreateFile(path)
	defer deferCloseFile(f) // defer will execute at the end of the function
	deferWriteFile(f)
}
