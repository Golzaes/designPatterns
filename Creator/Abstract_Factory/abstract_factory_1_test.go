package Abstract_Factory

import "testing"

func TestNewLunch(t *testing.T) {
	facory := NewLunch()

	// 主食
	rise := facory.CookStapleFood()
	rise.Cook()

	// 配菜
	beef := facory.CookDish()
	beef.Cook()
}
