package context

import (
	"github.com/free5gc/fsm"
)

const (
	Deregistered            fsm.StateType = "Deregistered"
	DeregistrationInitiated fsm.StateType = "DeregistrationInitiated"
	Authentication          fsm.StateType = "Authentication"
	SecurityMode            fsm.StateType = "SecurityMode"
	ContextSetup            fsm.StateType = "ContextSetup"
	Registered              fsm.StateType = "Registered"
)

type DummyAmfUe struct {
	State *fsm.State
}

func (ue *DummyAmfUe) init() {
	ue.State = fsm.NewState(Deregistered)
}

func NewDummyAmfUe() *DummyAmfUe {
	ue := DummyAmfUe{}
	ue.init()
	return &ue
}
