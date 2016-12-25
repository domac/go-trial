package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	WRITE_APPEND = 0
	WRITE_OVER   = 1
)

//删除文件
func removeFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func createFile(filepath string) (string, error) {
	finfo, err := os.Stat(filepath)
	if err == nil {
		if finfo.IsDir() {
			return filepath, errors.New("filepath is a dir")
		} else {
			return filepath, errors.New("filepath exists")
		}
	}
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("File path is not exist")
		return filepath, err
	}
	defer f.Close()
	return filepath, nil
}

//以追加的方式打开
func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

//以覆盖的方式打开
func openToOverwrite(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

//写入文件
func writeFile(filepath string, content []string, writeMode int) error {

	var f *os.File
	var err error

	if writeMode == WRITE_APPEND {
		f, err = openToAppend(filepath)
	} else {
		f, err = openToOverwrite(filepath)
	}

	if err != nil {
		return err
	}
	defer f.Close()

	for _, s := range content {
		fmt.Fprintln(f, s)
	}
	return nil
}

func main() {
	filepath := "/tmp/yy.txt"
	// err := removeFile(filepath)
	// if err != nil {
	// 	fmt.Printf("file remove fail: %s\n", filepath)
	// } else {
	// 	fmt.Printf("file remove success: %s \n", filepath)
	// }

	// path, err := createFile(filepath)
	// if err != nil {
	// 	fmt.Printf("file create fail: %s\n", err.Error())
	// } else {
	// 	fmt.Printf("file create success: %s \n", path)
	// }

	content := []string{
		"192.168.139.1",
		"192.168.139.2",
	}

	writeFile(filepath, content, WRITE_OVER)
}
