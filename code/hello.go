package main

import "fmt"

type Presenter struct {
	Name      string
	Employeer string
	Title     string
	Role      string
}

func main() {
	x := Presenter{}
	x.Name = "Sean Malloy"
	x.Employeer = "Kohl's Departments Stores"
	x.Title = "Software Engineer"
	x.Role = "Cloud Platform Automation(OpenShift, k8s, Vault)"
	fmt.Println(x.Name)
	fmt.Println(x.Employeer)
	fmt.Println(x.Title)
	fmt.Println(x.Role)
}
