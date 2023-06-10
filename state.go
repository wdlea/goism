package goism

type Action uint8
type StateID uint64

const (
	Action_Repeat Action = iota
	Action_TerminateMachine
	Action_Next
)

type State struct {
	ID StateID

	Call func() Action

	NextStates []StateID
}

func (s *State) Evaluate() (result Action) {
	return s.Call()
}

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
			cont = len(s.NextStates) > int(result)
			next = s.NextStates[int(result)]
		}
	}
	return
}
