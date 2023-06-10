package goism

type Action uint8
type StateID uint64

const (
	Action_Repeat Action = iota
	Action_TerminateMachine
	Action_Next
)

// A state is one "Node" in a flowchart
// representation of the machine. The
// NextStates are called based on the returned
// Action from the call minus Action_Next
// use the constant instead of the value to
// allow for more enum values to be added.
// For example returning Action_Next + 1
// would call the state with the ID in
// index 1 of the slice.
type State struct {
	ID StateID

	Call func() Action

	NextStates []StateID
}

// Evaluate runs the State and returns the action. NOTE that
// this will not advance or change the machine in any way.
func (s *State) Evaluate() (result Action) {
	return s.Call()
}

// SelectNext returns the next StateID or whether to
// continue running(cont) based on the action.
func (s *State) SelectNext(result Action) (cont bool, next StateID) {
	switch result {
	case Action_Repeat:
		{
			cont = true
			next = s.ID
			break
		}
	case Action_TerminateMachine:
		{
			cont = false
			break
		}
	default:
		{
			t_next := int(result) - int(Action_Next)
			cont = len(s.NextStates) > t_next
			next = s.NextStates[t_next]
		}
	}
	return
}
