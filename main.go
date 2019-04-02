package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// get url from flag
	url_name := flag.String("img", "", "URL of image to download")
	flag.Parse()

	var loc string
	if *url_name == "" {
		loc = "/System/Resources/Golang"
		file := filepath.Join(loc, "imgs.txt")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		contents := strings.Split(string(data), "\n")
		urls := contents[:len(contents)-1]

		// get the first url and move to bottom of file
		*url_name = urls[0]
		f, err := os.OpenFile(file, os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		lines := strings.Join(append(urls[1:], urls[0]), "\n")
		if _, err = f.WriteString(lines); err != nil {
			panic(err)
		}
		log.Println("File written successfully.")
	}

	filePath, err := GetImg(*url_name, loc)
	if err != nil {
		panic(err)
	}
	s := "gsettings set org.gnome.desktop.background picture-uri file://"
	args := strings.Split(s, " ")
	args[len(args)-1] += filePath
	cmd := exec.Command(args[0], args[1:]...)
	log.Printf("Setting desktop to %s", filePath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
