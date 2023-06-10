package goism

// A StateMachineSchema contains all the states a StateMachine can have
// and is used as a blueprint via the CreateInstance function.
// A StateMachineSchema cannot be run without first creating an instance.
type StateMachineSchema struct {
	states []*State
}

// Add state adds a state to the StateMachineSchema and returns its StateID
// with the intention that it is used to address the added state from other states.
func (schema *StateMachineSchema) AddState(s *State) (id StateID) {
	id = StateID(len(schema.states))

	schema.states = append(schema.states, s)
	return
}

// CreateInstance creates an instance of the StatemachineSchema
// if orphan is set to true, the resulting StateMachineInstance is an
// "orphan", which means that it has its own copy of the schema,
// allowing you to change it at runtime without other stateMachines
// being affected.
func (schema *StateMachineSchema) CreateInstance(orphan bool) StateMachineInstance {
	var parent *StateMachineSchema

	if orphan {
		schemaCopy := *schema
		parent = &schemaCopy
	} else {
		parent = schema
	}

	return StateMachineInstance{
		CurrentState: 0,
		Schema:       parent,
	}
}

// A StateMachineInstance is an instance of a StateMachineSchema
type StateMachineInstance struct {
	CurrentState StateID
	Schema       *StateMachineSchema
}

// EvaluateCurrent evaluates the current state once, without continuing
// to the next state, subsequent calls to this will advance the stateMachine.
// For continuious running use EvaluateRecursively. EvaluateCurrent evaluates
// a single "tick" of the state machine. When the stateMachine has no more states
// to advance to, stop will be true.
func (m *StateMachineInstance) EvaluateCurrent() (stop bool) {
	state := m.Schema.states[m.CurrentState]
	result := state.Evaluate()
	cont, next := state.SelectNext(result)

	if cont {
		m.CurrentState = next
	} else {
		m.CurrentState = 0
	}
	return cont
}

// EvaluateRecursively repeatedly calls
func (m *StateMachineInstance) EvaluateRecursively(stopCallResets bool) {
	for {
		stop := m.EvaluateCurrent()
		if stop {
			if !stopCallResets {
				return
			}
		}
	}
}
