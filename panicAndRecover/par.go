package par

import (
	"fmt"
	"runtime"
)

func foo() {
	defer fmt.Println("Exiting foo")
	panic("bar")
}

func PanicRecovery(err *error) {
	if r := recover(); r != nil {
		if _, ok := r.(runtime.Error); ok {
			*err = r.(error)
		} else {
			*err = r.(error)
		}
	}
}

type Foo struct {
	Is []int
}

func (fp *Foo) Panic() (err error) {
	defer PanicRecovery(&err)
	fp.Is[0] = 5
	return nil
}
func Run() {
	defer fmt.Println("Exiting main")
	fp := &Foo{}
	if err := fp.Panic(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//foo()
	fmt.Println("ok")
}

func PanicIf(cond bool, args ...interface{}) {
	if !cond {
		return
	}
	if len(args) == 0 {
		panic(fmt.Errorf("cond failed"))
	}

	format := args[0].(string)
	args = args[1:]
	err := fmt.Errorf(format, args...)
	panic(err)
}
