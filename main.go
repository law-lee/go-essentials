package main

import (
	errordemo "github.com/law-lee/go-essentials/error-demo"
	functiondemo "github.com/law-lee/go-essentials/function-demo"
	"github.com/law-lee/go-essentials/methods"
	"github.com/law-lee/go-essentials/defer-demo"
)

func main() {
	functiondemo.Run()
	methods.Run()
	errordemo.Run()
	deferdemo.Run()
}
