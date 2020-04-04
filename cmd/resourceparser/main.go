package main

import (
	"github.com/smm-goddess/resourceparser/resources/parser"
	"io/ioutil"
)

var resourcePath = "F:/AndroidStudioProjects/MyApplication2/app/build/outputs/apk/debug/app-debug/resources.arsc"

func main() {
	bs, _ := ioutil.ReadFile(resourcePath)
	parser.Parse(bs)
}
