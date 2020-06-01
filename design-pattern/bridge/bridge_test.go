package bridge_test

import (
	"github.com/go-study/src/design-pattern/bridge"
	"testing"
)

func TestBridge(t *testing.T) {
	redCircle := new(bridge.Circle)
	redCircle.Circle(1, 2, 3, new(bridge.RedCircle))
	redCircle.Draw()

	greenCircle := new(bridge.Circle)
	greenCircle.Circle(4, 5, 6, new(bridge.GreenCircle))
	greenCircle.Draw()
}
