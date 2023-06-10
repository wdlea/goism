package goism

import "testing"

func Test_all(t *testing.T) {
	schema := StateMachineSchema{}

	foo := 1

	stateIntial := State{
		Call: func() Action {
			foo = 69

			return Action_Next
		},
		NextStates: make([]StateID, 2),
	}
	stateSecond := State{
		Call: func() Action {
			foo = 7
			return Action_TerminateMachine
		},
		NextStates: make([]StateID, 0),
	}
	stateSkipped := State{
		Call: func() Action {
			foo = 8
			return Action_Next
		},
		NextStates: make([]StateID, 1),
	}

	schema.AddState(&stateIntial)

	secondStateID := schema.AddState(&stateSecond)
	skippedStateID := schema.AddState(&stateSkipped)

	stateIntial.NextStates[0] = secondStateID
	stateIntial.NextStates[1] = skippedStateID
	stateSkipped.NextStates[0] = stateIntial.ID

	instance := schema.CreateInstance(false)

	stop := instance.EvaluateCurrent()

	if foo != 69 {
		t.Fatalf("Invalid foo value, initial state not effective")
	}
	if stop {
		t.Fatalf("Machine stopped too soon")
	}

	stop = instance.EvaluateCurrent()

	if foo != 7 {
		t.Fatalf("Invalid foo value, second state not effective")
	}

	if !stop {
		t.Fatalf("State machine has no brakes")
	}

}
