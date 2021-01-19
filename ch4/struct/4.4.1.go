package main

import "fmt"

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 /100
}

func main() {
	//p := Point{1, 2}
	fmt.Println(Scale(Point{1, 2}, 5))  // "{5 10}"
}
