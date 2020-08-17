package typesystem

import "fmt"

type Speaker struct {
	Name string
	age  int
}

func (s *Speaker) Speak() {
	fmt.Printf("now %s is speaking", s.Name)
}

// Do<number> function is use case for example
func Do1() {
	las := &Speaker{
		Name: "Las",
		age:  20,
	}
	las.Speak()
}
