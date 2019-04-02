// +build linux darwin
package main

func GetCommand() string {
	return "gsettings set org.gnome.desktop.background picture-uri file://"
}

func Update() {}
