package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name     string
	Age      int
	JobTitle string
}

func main() {
	employee := &Employee{
		Name: "yanle",
		Age:  32,
	}

	employee.Name = "yanle"
}
