package deferdemo

import "fmt"

func logNum(i int) {
	fmt.Printf("Num %d\n", i)
}

func plusOne(i int) (result int) {
	// anonymous function must be called by adding ()
	defer func() { result++ }()

	// i is returned as result, which is updated by deferred function above
	// after execution of below return
	return i
}

func Run() {
	defer logNum(1)
	fmt.Println("First main statement")
	defer logNum(2)
	defer logNum(3)
	fmt.Println(plusOne(1))

	// defer creates a closure which only captures variable i by a reference. It doesn't capture the value of the variable.
	// go1.22修复了这个问题，每次循环迭代，都会创建一个新的循环变量 i，并用上一轮迭代结束时的值来初始化它。
	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Printf("value in defer is: %d\n", i)
		}()
	}

	// A closure might be slightly more expensive as it requires
	// allocating an object to collect all the variables captured by the closure.
	for i := 0; i < 2; i++ {
		defer func(i2 int) {
			fmt.Printf("value in defer with param is: %d\n", i2)
		}(i)
	}

	panic("panic occurred after defer")

	fmt.Println("Last main statement") // not printed

	// not deferred since execution flow never reaches this line
	defer logNum(4)
}
