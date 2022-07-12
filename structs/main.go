package main

import (
	"fmt"
	"math"
	"strconv"
)

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

type Square struct {
	x float64
}

type Triangle struct {
	x, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (s Square) area() float64 {
	return s.x * s.x
}

func (t Triangle) area() float64 {
	return t.x * t.h / 2
}

func getArea(s Shape) float64 {
	return s.area()
}

type Person struct {
	firstName, lastName string
	age                 int
}

// value receiver
func (p Person) greet() string {
	return "Hiya! " + p.firstName + " " + p.lastName + " here and I am " + strconv.Itoa(p.age) + " years old."
}

// pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{firstName: "John", lastName: "Doe", age: 30}
	person2 := Person{"Jane", "Doe", 25}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Print(person1.age)

	// interface
	circle := Circle{r: 5}
	square := Square{x: 4}
	triangle := Triangle{x: 3, h: 4}

	fmt.Printf("circle => %f\n", getArea(circle))
	fmt.Printf("square => %f\n", getArea(square))
	fmt.Printf("triangle => %f\n", getArea(triangle))
}
