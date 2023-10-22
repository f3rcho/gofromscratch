package models

type Man struct {
	age    int
	height float32
	weight float32
	breath bool
	think  bool
	eat    bool
	alive  bool
}

func (m *Man) Breath() {
	m.breath = true
	m.think = false
	m.eat = false
	m.alive = true
}

func (m *Man) Eat()        { m.eat = true }
func (m *Man) Think()      { m.think = true }
func (m *Man) Sex() string { return "Man" }
