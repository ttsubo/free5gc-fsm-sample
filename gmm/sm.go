package gmm

import (
	"fmt"
	"github.com/free5gc/fsm"
)

func DeRegistered(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		fmt.Printf("Receiving Event: [%+v] on DeRegistered\n", event)
	case GmmMessageEvent:
		fmt.Printf("Receiving Event: [%+v] on DeRegistered\n", event)
		fmt.Println("... Send Event: StartAuthEvent")
		GmmFSM.SendEvent(state, StartAuthEvent, fsm.ArgsType{})
	case StartAuthEvent:
		fmt.Printf("Receiving Event: [%+v] on DeRegistered\n", event)
	case fsm.ExitEvent:
		fmt.Printf("Receiving Event: [%+v] on DeRegistered\n", event)
	default:
		fmt.Println("Unknown event [%+v]", event)
	}
}

func Authentication(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		fmt.Printf("Receiving Event: [%+v] on Authentication\n", event)
		fallthrough
	case AuthRestartEvent:
		fmt.Printf("(Receiving Event: AuthRestartEvent on Authentication)\n")
	case GmmMessageEvent:
		fmt.Printf("Receiving Event: [%+v] on Authentication\n", event)
		fmt.Println("... Send Event: AuthSuccessEvent")
		GmmFSM.SendEvent(state, AuthSuccessEvent, fsm.ArgsType{})
	case AuthSuccessEvent:
		fmt.Printf("Receiving Event: [%+v] on Authentication\n", event)
	case fsm.ExitEvent:
		fmt.Printf("Receiving Event: [%+v] on Authentication\n", event)
	default:
		fmt.Println("Unknown event [%+v]", event)
	}
}

func SecurityMode(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		fmt.Printf("Receiving Event: [%+v] on SecurityMode\n", event)
	case GmmMessageEvent:
		fmt.Printf("Receiving Event: [%+v] on SecurityMode\n", event)
		fmt.Println("... Send Event: SecurityModeSuccessEvent")
		GmmFSM.SendEvent(state, SecurityModeSuccessEvent, fsm.ArgsType{})
	case SecurityModeSuccessEvent:
		fmt.Printf("Receiving Event: [%+v] on SecurityMode\n", event)
	case fsm.ExitEvent:
		fmt.Printf("Receiving Event: [%+v] on SecurityMode\n", event)
	default:
		fmt.Println("Unknown event [%+v]", event)
	}
}

func ContextSetup(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		fmt.Printf("Receiving Event: [%+v] on ContextSetup\n", event)
	case GmmMessageEvent:
		fmt.Printf("Receiving Event: [%+v] on ContextSetup\n", event)
		fmt.Println("... Send Event: ContextSetupSuccessEvent")
		GmmFSM.SendEvent(state, ContextSetupSuccessEvent, fsm.ArgsType{})
	case ContextSetupSuccessEvent:
		fmt.Printf("Receiving Event: [%+v] on ContextSetup\n", event)
	case fsm.ExitEvent:
		fmt.Printf("Receiving Event: [%+v] on ContextSetup\n", event)
	default:
		fmt.Printf("Unknown event [%+v]", event)
	}
}

func Registered(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		fmt.Printf("Receiving Event: [%+v] on Registered\n", event)
	default:
		fmt.Printf("Unknown event [%+v]", event)
	}
}
