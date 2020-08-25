package bridge_test

import (
	"github.com/go-study/src/design-pattern/bridge"
	"testing"
)

func TestBridge(t *testing.T) {
	var shape bridge.Shape
	shape = bridge.NewCircle(1, 2, 3)

	// RedCircle
	shape.SetDrawAPI(new(bridge.RedCircle))
	shape.Draw()

	// GreenCircle
	shape.SetDrawAPI(new(bridge.GreenCircle))
	shape.Draw()
}
