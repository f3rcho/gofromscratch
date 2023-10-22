package models

type Woman struct {
	Man
	isMother bool
}

func (w *Woman) Breath() {
	w.breath = true
	w.think = false
	w.eat = false
	w.alive = true
	w.isMother = true
}

func (w *Woman) Eat()        { w.eat = true }
func (w *Woman) Think()      { w.think = true }
func (w *Woman) Sex() string { return "Woman" }
