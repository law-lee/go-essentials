package main

import (
	deferdemo "github.com/law-lee/go-essentials/defer-demo"
	errordemo "github.com/law-lee/go-essentials/error-demo"
	functiondemo "github.com/law-lee/go-essentials/function-demo"
	"github.com/law-lee/go-essentials/methods"
	par "github.com/law-lee/go-essentials/panicAndRecover"
)

func main() {
	functiondemo.Run()
	methods.Run()
	errordemo.Run()
	deferdemo.Run()
	par.Run()
}
