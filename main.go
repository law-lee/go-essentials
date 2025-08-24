package main

import (
	"github.com/law-lee/go-essentials/base64encoding"
	"github.com/law-lee/go-essentials/channels"
	"github.com/law-lee/go-essentials/commands"
	"github.com/law-lee/go-essentials/concurrency"
	"github.com/law-lee/go-essentials/csvdemo"
	"github.com/law-lee/go-essentials/datetime"
	deferdemo "github.com/law-lee/go-essentials/defer-demo"
	errordemo "github.com/law-lee/go-essentials/error-demo"
	"github.com/law-lee/go-essentials/fileio"
	functiondemo "github.com/law-lee/go-essentials/function-demo"
	"github.com/law-lee/go-essentials/jsondemo"
	"github.com/law-lee/go-essentials/methods"
	mutexdemo "github.com/law-lee/go-essentials/mutex-demo"
	par "github.com/law-lee/go-essentials/panicAndRecover"
	"github.com/law-lee/go-essentials/xmldemo"
	"github.com/law-lee/go-essentials/reflection"
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
	fileio.Run()
	datetime.Run()
	commands.Run()
	base64encoding.Run()
	jsondemo.Run()
	xmldemo.Run()
	xmldemo.Run2()
	xmldemo.Run3()
	csvdemo.Run()
	reflection.Run()
}
