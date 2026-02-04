package core

import "fmt"

type State string // Defining State as a string type

const ( // Defining possible states
	Init     State = "init"
	Running  State = "running"
	Draining State = "draining"
	Stopped  State = "stopped"
)

var validTransitions = map[State][]State{ // Defining valid state transitions
	Init:     {Running},
	Running:  {Draining},
	Draining: {Stopped},
}

func CanTransition(from, to State) bool { // Check if a transition is valid
	for _, s := range validTransitions[from] {
		if s == to {
			return true
		}
	}
	return false
}

func TransitionError(from, to State) error { // Return an error for invalid transitions
	return fmt.Errorf("invalid transition: %s → %s", from, to)
}
