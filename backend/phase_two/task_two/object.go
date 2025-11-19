package task_two

import (
	"math"
	"strconv"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func area(s Shape) float64 {
	return s.Area()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeId int
	Person
}

func (e *Employee) GetInfo() string {
	return "EmployeeId:" + strconv.Itoa(e.EmployeeId) + " Name:" + e.Name +
		" Age:" + strconv.Itoa(e.Age)
}
