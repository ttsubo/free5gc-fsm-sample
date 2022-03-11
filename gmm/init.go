package gmm

import (
	"fmt"
	"github.com/free5gc/fsm"
	"github.com/ttsubo/free5gc-fsm-sample/context"
)

const (
	GmmMessageEvent          fsm.EventType = "Gmm Message"
	StartAuthEvent           fsm.EventType = "Start Authentication"
	AuthSuccessEvent         fsm.EventType = "Authentication Success"
	AuthRestartEvent         fsm.EventType = "Authentication Restart"
	SecurityModeSuccessEvent fsm.EventType = "SecurityMode Success"
	ContextSetupSuccessEvent fsm.EventType = "ContextSetup Success"
)

var transitions = fsm.Transitions{
	{Event: GmmMessageEvent, From: context.Deregistered, To: context.Deregistered},
	{Event: GmmMessageEvent, From: context.Authentication, To: context.Authentication},
	{Event: GmmMessageEvent, From: context.SecurityMode, To: context.SecurityMode},
	{Event: GmmMessageEvent, From: context.ContextSetup, To: context.ContextSetup},
	{Event: GmmMessageEvent, From: context.Registered, To: context.Registered},
	{Event: StartAuthEvent, From: context.Deregistered, To: context.Authentication},
	{Event: AuthSuccessEvent, From: context.Authentication, To: context.SecurityMode},
	{Event: SecurityModeSuccessEvent, From: context.SecurityMode, To: context.ContextSetup},
	{Event: ContextSetupSuccessEvent, From: context.ContextSetup, To: context.Registered},
}

var callbacks = fsm.Callbacks{
	context.Deregistered:   DeRegistered,
	context.Authentication: Authentication,
	context.SecurityMode:   SecurityMode,
	context.ContextSetup:   ContextSetup,
	context.Registered:     Registered,
}

var GmmFSM *fsm.FSM

func init() {
	if f, err := fsm.NewFSM(transitions, callbacks); err != nil {
		fmt.Printf("Initialize Gmm FSM Error: %+v", err)
	} else {
		GmmFSM = f
	}
}
