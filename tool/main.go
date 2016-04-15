package main

import (
	"io/ioutil"
	"os"
)

func main() {
	links()
	code, _ := tostring()
	ioutil.WriteFile("../api.go", []byte(code), os.ModePerm)
}
