package main

import (
	"encoding/gob"
	"fmt"
	"github.com/mreiferson/go-snappystream"
	"os"
)

type FileCompress struct {
	Classifier interface{}
}

func (t *FileCompress) SaveFile(filename string) error {
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fi.Close()

	fs := snappystream.NewBufferedWriter(fi)
	encoder := gob.NewEncoder(fs)
	err = encoder.Encode(t.Classifier)
	if err != nil {
		return err
	}
	err = fs.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	obj := FileCompress{Classifier: "it works"}
	err := obj.SaveFile("test.sz")
	if err != nil {
		fmt.Println(err)
	}
}
