package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// get url from flag
	url_name := flag.String("img", "", "URL of image to download")
	flag.Parse()

	if *url_name == "" {
		// get random url from file
		data, err := ioutil.ReadFile("imgs.txt")
		if err != nil {
			fmt.Println("Error reading file", err)
		}
		contents := strings.Split(string(data), "\n")
		urls := contents[:len(contents)-1]

		// get the first url and move to bottom of file
		*url_name = urls[0]
		f, err := os.OpenFile("imgs.txt", os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		lines := strings.Join(append(urls[1:], urls[0]), "\n")
		if _, err = f.WriteString(lines); err != nil {
			panic(err)
		}
		fmt.Println("File written successfully.")
	}

	filePath, err := GetImg(*url_name)
	if err != nil {
		panic(err)
	}
	s := "gsettings set org.gnome.desktop.background picture-uri file://"
	args := strings.Split(s, " ")
	args[len(args)-1] += filePath
	cmd := exec.Command(args[0], args[1:]...)
	fmt.Printf("Setting desktop to %s", filePath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
