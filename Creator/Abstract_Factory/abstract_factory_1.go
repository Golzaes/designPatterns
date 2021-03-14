package Abstract_Factory

import "fmt"

// 自顶向下
// Food
type Food interface {
	CookStapleFood() Menu
	CookDish() Menu
}

type Menu interface {
	Cook()
}

type Lunch struct {
}

func NewLunch() Food{
	return &Lunch{}
}

func (l *Lunch) CookStapleFood()  Menu{
	return &Rise{}
}
func (l *Lunch) CookDish()  Menu{
	return &Beef{}
}

type Rise struct {
}

func (r *Rise) Cook()  {
	fmt.Println("Create Rise~")
}
type Beef struct {
}

func (b *Beef) Cook()  {
	fmt.Println("Create Beef~")
}
