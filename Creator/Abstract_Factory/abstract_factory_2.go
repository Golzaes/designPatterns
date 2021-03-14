package Abstract_Factory

import "fmt"

// 定义食物种类
type Foods interface {
	CreateStapleFoods() Cook
	CreateDishFoods() Cook
}

// 定义制作食物
type Cook interface {
	Cooks()
}

type LunchFood struct {
}

func NewLunchFood() Foods {
	return &LunchFood{}
}

func (l LunchFood) CreateStapleFoods() Cook {
	return &Bun{}
}
func (l LunchFood) CreateDishFoods() Cook {
	return &Cack{}
}

type Bun struct {
}

func (b *Bun) Cooks() {
	fmt.Println("Cooks Bun~")
}

type Cack struct {
}

func (c *Cack) Cooks() {
	fmt.Println("Cooks Cack~")
}
