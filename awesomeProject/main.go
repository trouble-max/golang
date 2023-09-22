package main

import (
	"awesomeProject/accountant"
	"awesomeProject/engineer"
	"awesomeProject/janitor"
	"awesomeProject/manager"
	"awesomeProject/salesman"
	"fmt"
)

func main() {
	a1 := accountant.Accountant{}
	a1.SetSalary(500.00)
	a1.SetPosition("Accountant")
	a1.SetAddress("Address1")

	e1 := engineer.Engineer{}
	e1.SetSalary(1100.00)
	e1.SetPosition("Engineer")
	e1.SetAddress("Address2")

	j1 := janitor.Janitor{}
	j1.SetSalary(300.00)
	j1.SetPosition("Janitor")
	j1.SetAddress("Address3")

	m1 := manager.Manager{}
	m1.SetSalary(900.00)
	m1.SetPosition("Manager")
	m1.SetAddress("Address4")

	s1 := salesman.Salesman{}
	s1.SetSalary(800.00)
	s1.SetPosition("Salesman")
	s1.SetAddress("Address5")

	fmt.Println(a1)
	fmt.Println(e1)
	fmt.Println(j1)
	fmt.Println(m1)
	fmt.Println(s1)
}
