package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 24})
	fmt.Println(Person{name: "Alice",age: 25})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age:40})
	fmt.Println(newPerson("Jon"))

	s := Person{name:"Sean", age: 24}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}