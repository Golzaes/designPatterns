package Abstract_Factory

import "testing"

func TestNewLunchFood(t *testing.T) {
	facory := NewLunchFood()

	// 主食
	rise := facory.CreateStapleFoods()
	rise.Cooks()

	// 配菜
	beef := facory.CreateDishFoods()
	beef.Cooks()
}