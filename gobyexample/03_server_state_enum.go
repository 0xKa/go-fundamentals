package gobyexample

import "fmt"

type serverState int

const (
	// iota generates sequential values for the constants
	serverStateIdle serverState = iota
	serverStateStopped
	serverStateError
	serverStateRunning
	serverStateConnected
	serverStateRetrying
)

var serverStateNames = map[serverState]string{
	serverStateIdle:      "Idle",
	serverStateStopped:   "Stopped",
	serverStateError:     "Error",
	serverStateRunning:   "Running",
	serverStateConnected: "Connected",
	serverStateRetrying:  "Retrying",
}

func (s serverState) String() string {
	if name, ok := serverStateNames[s]; ok {
		return name
	}
	return fmt.Sprintf("ServerState(%d)", s)
}

func transitionServerState(s serverState) serverState {
	switch s {
	case serverStateIdle:
		return serverStateConnected
	case serverStateConnected, serverStateRetrying:
		return serverStateIdle
	case serverStateError:
		return serverStateError
	case serverStateRunning:
		return serverStateStopped
	case serverStateStopped:
		return serverStateRunning
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func ShowServerStateEnum() {
	var s serverState
	for s = serverStateIdle; s <= serverStateRetrying; s++ {
		fmt.Println(s)
	}

	fmt.Println("\nTransitioning states:")
	ns := transitionServerState(serverStateIdle)
	fmt.Println(ns)

	ns2 := transitionServerState(ns)
	fmt.Println(ns2)
}
