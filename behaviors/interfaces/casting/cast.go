// Will this code compile? panic?
package main

type IntOne interface {
	Name() string
}

type IntTwo interface {
	Age() int
}

type Person struct {
	name string
	age  int
}

func (p Person) Name() string {
	return p.name
}

func (p Person) Age() int {
	return p.age
}

func main() {

	p := Person{
		name: "Samit",
		age:  20,
	}

	Hello(p)

}

func Hello(one IntOne) {
	println("got one: ", one.Name())

	two := one.(IntTwo)
	println("got two: ", two.Age())

}
