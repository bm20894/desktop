//+build windows
package main

import "os/exec"

func GetCommand() string {
	return "reg add \"HKEY_CURRENT_USER\\Control Panel\\Desktop\" /v Wallpaper /t REG_SZ /d  wallpaper_path /f"
}

func Update() {
	cmd := exec.Command("RUNDLL32.EXE", "user32.dll,UpdatePerUserSystemParameters")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
