package build

import "testing"

func TestFieldsOf(t *testing.T) {
	player := InitPlayer()
	t.Log(player.Name)
}
