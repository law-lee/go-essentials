package main

import (
	"github.com/law-lee/go-essentials/channels"
	"github.com/law-lee/go-essentials/concurrency"
	deferdemo "github.com/law-lee/go-essentials/defer-demo"
	errordemo "github.com/law-lee/go-essentials/error-demo"
	functiondemo "github.com/law-lee/go-essentials/function-demo"
	"github.com/law-lee/go-essentials/methods"
	mutexdemo "github.com/law-lee/go-essentials/mutex-demo"
	par "github.com/law-lee/go-essentials/panicAndRecover"
)

func main() {
	functiondemo.Run()
	methods.Run()
	errordemo.Run()
	deferdemo.Run()
	par.Run()
	concurrency.Run()
	concurrency.Run2()
	channels.Run()
	mutexdemo.Run()
}
