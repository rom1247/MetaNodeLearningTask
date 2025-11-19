package task_two

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {

	circle := &Circle{5}
	rectangle := &Rectangle{5, 10}

	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())
	fmt.Println(circle.Perimeter())
	fmt.Println(rectangle.Perimeter())

	fmt.Println(area(circle))
	fmt.Println(area(rectangle))

	employee := Employee{1, Person{"张三", 18}}
	fmt.Println(employee.GetInfo())
}
