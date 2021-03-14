package Factory

import "fmt"

type Cook interface {
	CookFood()
}

type ChefMan struct {
	Id int
}
type ChefWoman struct {
	Id int
}

func (m ChefMan) CookFood() {
	fmt.Println("ChefMan Cooking~")
}

func (w ChefWoman) CookFood() {
	fmt.Println("ChefWoman Cooking~")
}

func NewDinner(i int) Cook {
	switch i {
	case 1:
		Cook.CookFood(ChefMan{i})
	case 2:
		Cook.CookFood(ChefWoman{i})
	}
	return nil
}
