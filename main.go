package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	url_name := flag.String("img", "http://paperlief.com/images/hd-fall-nature-wallpapers-wallpaper-1.jpg", "URL of image to download")
	flag.Parse()
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
