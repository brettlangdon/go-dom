package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func copyFile(fname string) error {
	base := path.Base(fname)
	in, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(base, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	b, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	_, err = out.Write(b)
	return err
}

func main() {
	fnames, err := filepath.Glob("generate/static/*.go")
	if err != nil {
		panic(err)
	}
	for _, n := range fnames {
		err = copyFile(n)
		if err != nil {
			panic(err)
		}
	}

	generator, err := NewGenerator("generate/api.min.json")
	if err != nil {
		panic(err)
	}

	err = generator.Generate()
	if err != nil {
		panic(err)
	}
}
