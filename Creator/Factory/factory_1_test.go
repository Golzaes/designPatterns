package Factory

import "testing"

func TestNewStore(t *testing.T) {
	NewStore("w")
	NewStore("e")
}
