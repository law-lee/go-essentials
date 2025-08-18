package main

import (
	errordemo "github.com/law-lee/go-essentials/error-demo"
	functiondemo "github.com/law-lee/go-essentials/function-demo"
	"github.com/law-lee/go-essentials/methods"
)

func main() {
	functiondemo.Run()
	methods.Run()
	errordemo.Run()
	errordemo.PanicIf(false, "no panic happens")
}
