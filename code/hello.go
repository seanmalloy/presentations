package main

import "fmt"

type Presenter struct {
	Name      string
	Employeer string
	Title     string
}

func main() {
	x := Presenter{}
	x.Name = "Sean Malloy"
	x.Employeer = "Kohl's Departments Stores"
	x.Title = "Platform Engineer"
	fmt.Println(x.Name)
	fmt.Println(x.Employeer)
	fmt.Println(x.Title)
}
