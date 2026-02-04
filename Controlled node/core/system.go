package core

import (
	"log"
	"sync"
)

type System struct { // System struct to hold state and mode
	mu    sync.Mutex
	state State
	mode  string
}

func NewSystem() *System { // Constructor for System
	return &System{
		state: Init,
		mode:  "normal",
	}
}

func (s *System) GetState() State { // Getter for current state
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.state
}

func (s *System) GetMode() string { // Getter for current mode
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.mode
}

func (s *System) Transition(to State) error { // Method to transition state
	s.mu.Lock()
	defer s.mu.Unlock()

	if !CanTransition(s.state, to) { // Check if transition is valid
		log.Println(TransitionError(s.state, to))
		return TransitionError(s.state, to)
	}

	log.Printf("STATE CHANGE: %s → %s\n", s.state, to) // Log state change
	s.state = to
	return nil
}

func (s *System) SetMode(mode string) { // Method to set mode
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Println("MODE CHANGE:", mode)
	s.mode = mode
}
