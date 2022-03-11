package main

import (
	"fmt"
	"github.com/free5gc/fsm"
	"github.com/ttsubo/free5gc-fsm-sample/context"
	"github.com/ttsubo/free5gc-fsm-sample/gmm"
	"time"
)

func main() {
	ue := context.NewDummyAmfUe()

	time.Sleep(2 * time.Second)
	fmt.Println("... Send Event: [GmmMessageEvent], after receiving [NAS-PDU: Registration request]")
	gmm.GmmFSM.SendEvent(ue.State, gmm.GmmMessageEvent, fsm.ArgsType{})
	fmt.Println("")

	time.Sleep(5 * time.Second)
	fmt.Println("... Send Event: [GmmMessageEvent], after receiving [NAS-PDU: Authentication Response]")
	gmm.GmmFSM.SendEvent(ue.State, gmm.GmmMessageEvent, fsm.ArgsType{})
	fmt.Println("")

	time.Sleep(5 * time.Second)
	fmt.Println("... Send Event: [GmmMessageEvent], after receiving [NAS-PDU: Security mode complete]")
	gmm.GmmFSM.SendEvent(ue.State, gmm.GmmMessageEvent, fsm.ArgsType{})
	fmt.Println("")

	time.Sleep(5 * time.Second)
	fmt.Println("... Send Event: [GmmMessageEvent], after receiving [NAS-PDU: Registration complete]")
	gmm.GmmFSM.SendEvent(ue.State, gmm.GmmMessageEvent, fsm.ArgsType{})
}
