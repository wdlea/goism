# goism
***GO**lang **I**nfinite **S**tate **M**achine*

Goism is an implementation of the Infinite State Machine in go. It has support for runtime modification. An infinite state machine is defined as

>  An Infinite State machine is a state machine with an infinite amount of posible states. This can occur when new states are added by other states. An example of this is a game of chess, each move adds more possible moves.

## Example Code

*example.go*

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


## Infinite States

Runtime modification of state machines can be performed via editing the schema using the pointed value, creating an orphaned instance using

    schema.CreateInstance(true)

will make an instance with its own copy of the schema so it may be safely modified.