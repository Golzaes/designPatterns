package Factory

import "fmt"

type Store interface {
	Buy()
}

// West
type West struct {
	article string
}

func (w *West) Buy() {
	fmt.Println("In West Market")
}

type East struct {
	article string
}

func (e *East) Buy() {
	fmt.Println("In East Market")
}

// NewStore
func NewStore(name string) Store {
	switch name {
	case "w":
		return &West{}
	case "e":
		return &East{}
	default:
		fmt.Println(1)
	}
	return nil
}
