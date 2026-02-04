package observability

import "controlled-node/core"

type Status struct { // Struct to hold system status
	State string `json:"state"`
	Mode  string `json:"mode"`
}

func Snapshot(sys *core.System) Status { // Create a snapshot of the system status
	return Status{
		State: string(sys.GetState()),
		Mode:  sys.GetMode(),
	}
}
