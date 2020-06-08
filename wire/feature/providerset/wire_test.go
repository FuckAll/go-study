package build

import "testing"

func TestError(t *testing.T) {
	endingA, err := InitProviderSetEndingA("jetty", "mario")
	if err != nil {
		t.Error(err)
		return
	}
	endingA.Appear()
}
