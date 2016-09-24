package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println(resp.Status)
	println(string(resp_body))
	return nil

}

func main() {
	targetUrl := "xxx"
	filename := "/tmp/test.txt"
	postFile(filename, targetUrl)
}
