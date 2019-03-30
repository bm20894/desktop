# Desktop Image Tool
This is a simple tool written in Go that allows you to change your desktop background to an image using a URL.

## Installation
Install this project with `$ go get github.com/bm20894/desktop`. The `desktop` executable will be installed to `$GOPATH/bin/desktop`.

## Usage
Running the `desktop` executable with no arguments will use a default picture of `http://paperlief.com/images/hd-fall-nature-wallpapers-wallpaper-1.jpg`, a nice image of fall trees. You can specify an image to download with `$ desktop -img <url-of-image>` where `<url-of-image>` is the URL of the image you wish to download. This URL must be the address of an image file (ends with image file extension).

## Development
This project currently works on unix-like systems, support for Windows systems may come soon.
