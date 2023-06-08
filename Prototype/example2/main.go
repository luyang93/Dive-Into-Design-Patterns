package main

import "fmt"

func main() {
	circle := &Circle{Shape{10, 20, "red"}, 15}
	anotherCircle := circle.Clone()

	// 创建一个矩形
	rectangle := &Rectangle{Shape{30, 40, "blue"}, 50, 60}
	anotherRectangle := rectangle.Clone()

	// 打印原始和复制的对象
	fmt.Printf("Circle: %+v\n", circle)
	fmt.Printf("Cloned Circle: %+v\n", anotherCircle)
	fmt.Printf("Rectangle: %+v\n", rectangle)
	fmt.Printf("Cloned Rectangle: %+v\n", anotherRectangle)
}
