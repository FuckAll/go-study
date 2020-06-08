package build

import "testing"

func TestError(t *testing.T) {
	mission, err := InitMissionParam("jetty", "mario")
	if err != nil {
		t.Error(err)
		return
	}
	mission.Start()
}
