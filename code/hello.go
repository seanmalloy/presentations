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
	x.Employeer = "Red Hat"
	x.Title = "Technical Account Manager"
	fmt.Println(x.Name)
	fmt.Println(x.Employeer)
	fmt.Println(x.Title)
}
