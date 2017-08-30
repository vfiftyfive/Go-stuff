package main

import "fmt"

//Mutatable struct
type Mutatable struct {
	a int
	b int
}

//StayTheSame func
func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

//Mutate func
func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

func main() {
	m := Mutatable{0, 0}
	fmt.Println(m)
	m.StayTheSame()
	fmt.Println(m)
	m.Mutate()
	fmt.Println(m)
}
