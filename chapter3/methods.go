package chapter3

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perim() int {
	return 2 * r.width + 2 * r.height
}

func Methods() {
	rect:= rectangle{width:10, height:5}

	fmt.Println(rect)

	fmt.Println("area: ", rect.area(), "perim: ", rect.perim())
}
