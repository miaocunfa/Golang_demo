package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee {
	var e *Employee

	return e
}

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000  // 扣钱
	position := &dilbert.Position
	*position = "Senior " + *position // 升职

	fmt.Printf("%q", dilbert)
}

