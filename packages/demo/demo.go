package main

import "log"

type Person struct {
	Name string
}

func (l *Person) clone() *Person {
	nl := *l
	return &nl
}

func main() {
	person := &Person{"yanle"}

	person2 := person.clone()

	if person == person2 {
		log.Println("等于")
	} else {
		log.Println("不等于")
	}

	log.Println("person: ", person.Name)
	log.Println("person2: ", person2.Name)
}
