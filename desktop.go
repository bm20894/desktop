package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetImg(url_name string) (string, error) {
	var fileName string

	buildFileName(url_name, &fileName)
	file := createFile(&fileName)

	err := putFile(file, httpClient(), url_name)
	return fileName, err
}

func buildFileName(url_name string, fileName *string) {
	fileUrl, err := url.Parse(url_name)
	if err != nil {
		panic(err)
	}
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	*fileName = segments[len(segments)-1]
}

func createFile(fileName *string) *os.File {
	*fileName = filepath.Join("/System/Resources/Golang", *fileName)
	file, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return &client
}

func putFile(file *os.File, client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Printf("Just downloaded a file %s with size %v\n", file.Name(), size)

	return nil

}