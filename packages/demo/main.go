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
	person := Person{
		Name: "John",
		Age:  30,
	}

	employee := Employee(person)

}
