package build

import "testing"

func TestBuild(t *testing.T) {
	mission := InitMission("mario")
	mission.Start()

	mission = InitMissionParam("jetty", "mario")
	mission.Start()
}
