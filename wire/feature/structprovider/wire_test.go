package build

import "testing"

func TestStructProvider(t *testing.T) {
	endingA, err := InitProviderSetEndingA("jetty", "mario")
	if err != nil {
		t.Error(err)
		return
	}
	endingA.Appear()
}
