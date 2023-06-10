package main

import (
	"math/rand"

	"github.com/wdlea/goism"
)

const (
	Action_yes goism.Action = iota + goism.Action_Next
	Action_no               = iota + goism.Action_Next
)

func main() {
	schema := goism.StateMachineSchema{}

	initialState := goism.State{
		Call: func() goism.Action {
			println("Hello")

			coinFlip := rand.Int() % 2 // either 0 or 1

			if coinFlip == 0 {
				return Action_yes
			} else {
				return Action_no
			}
		},
		NextStates: make([]goism.StateID, 2), //make some room for next states
	}
	schema.AddState(&initialState) //0

	yesState := goism.State{
		Call: func() goism.Action {
			println("World")
			return goism.Action_TerminateMachine
		},
	}
	noState := goism.State{
		Call: func() goism.Action {
			println("There")
			return goism.Action_TerminateMachine
		},
	}

	yesStateID := schema.AddState(&yesState)
	noStateID := schema.AddState(&noState)

	initialState.NextStates[0] = yesStateID
	initialState.NextStates[1] = noStateID

	instance := schema.CreateInstance(false)

	instance.EvaluateRecursively(false)

	//Will print either
	// Hello\nThere
	// Hello\nWorld
}
