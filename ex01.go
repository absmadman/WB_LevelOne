package main

import "fmt"

type Human struct {
	name    string
	surName string
	gender  string
	age     uint8
}

func NewHuman(name string, surName string, gender string, age uint8) *Human {
	return &Human{
		name:    name,
		surName: surName,
		gender:  gender,
		age:     age,
	}
}

func (h Human) Speak() {
	fmt.Println("I am a human!")
}

func (h Human) Stats() {
	fmt.Println(h.name)
	fmt.Println(h.surName)
	fmt.Println(h.gender)
	fmt.Println(h.age)
}

type Action struct {
	height int
	Human
}

func NewAction(h Human, heig int) *Action {
	return &Action{
		height: heig,
		Human:  h,
	}
}

func main() {
	hum := NewHuman("Portyanaya", "Golova", "male", 23)

	hum.Speak()
	hum.Stats()
	act := NewAction(*hum, 176)

	act.Speak()
	act.Stats()
}
