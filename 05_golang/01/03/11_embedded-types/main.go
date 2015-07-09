package main

import "fmt"

type Person struct {
	Name string
}

type AndroidOne struct {
	Person Person
	Model  string
}

type AndroidTwo struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	a := new(AndroidOne)
	a.Person.Name = "Todd"
	a.Person.Talk()

	b := new(AndroidTwo)
	b.Name = "Todd"
	b.Talk()
}
