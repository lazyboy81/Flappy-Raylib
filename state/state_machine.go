package state

import "sync"

var gStateMachine *StateMachine
var once sync.Once

type StateFn func() State

type StateMachine struct {
	states  map[string]StateFn
	current State
}

func NewStateMachine(states map[string]StateFn) *StateMachine {
	once.Do(func() {
		gStateMachine = &StateMachine{
			states:  states,
			current: &emptyState{},
		}
	})

	return gStateMachine
}

func (s *StateMachine) Change(stateName string, params map[string]any) {
	s.current.Exit()
	s.current = s.states[stateName]() // call the create function
	s.current.Enter(params)
}

func (s *StateMachine) Update(dt float32) {
	s.current.Update(dt)
}

func (s *StateMachine) Render() {
	s.current.Render()
}
