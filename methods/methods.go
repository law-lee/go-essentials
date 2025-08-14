package methods

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintFullNameValue() {
	fmt.Printf("PrintFullNameValue:   address of p is %p\n", &p)
}

func (p *Person) PrintFullNamePointer() {
	fmt.Printf("PrintFullNamePointer: p is            %p\n", p)
}

type Foo struct {
	Str string
}

// the address of f receiver is different that the address of fv value.
// This indicates it's a different value i.e. f is a copy of fv that Go compiler created implicilty
// changing f inside Print value receiver doesn't change the fv value

func (f Foo) Print() {
	fmt.Printf("Value of Str inside Print: '%s'\n", f.Str)
	fmt.Printf("Address of &f  inside Print: 0x%p\n", &f)
	f.Str = "changed"
}

func Run() {
	p := Person{FirstName: "John", LastName: "Doe"}
	fmt.Printf("address of p:                         %p\n", &p)
	p.PrintFullNameValue()
	p.PrintFullNamePointer()

	pp := &p
	fmt.Printf("address of pp:                         %p\n", pp)
	pp.PrintFullNameValue()
	pp.PrintFullNamePointer()
	fv := Foo{
		Str: "Foo",
	}

	fmt.Printf("Address of &fv before Print: 0x%p\n", &fv)
	fv.Print()
	fmt.Printf("Value of Str after Print: '%s'\n", fv.Str)
}
