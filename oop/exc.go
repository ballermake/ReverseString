package oop

//Encapsulation
type User struct {
	Name string //Public
	age  int    //Private
}

//Inheritance
type Animal struct {
	Name string
}

type Dog struct {
	Animal //mewarisi properties yang ada di struct Animal
	Breed  string
}

//Polymorphism
type Shape interface {
	Area()
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

//Abstraction
type Vehicle interface {
	StartEngine()
}

func StartEngine() string {
	return "Engine Start!"
}
