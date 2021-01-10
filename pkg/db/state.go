package db

// State defines the status of the database Connection
type State string

const (
	// StateCreated means Connection object is created
	StateCreated State = "created"
	// StateOpening means the Connection object start to connect to the database
	StateOpening State = "opening"
	// StateConnected means the Connection object is connection to the database
	StateConnected State = "connected"
	// StateClosing means the Connection object starts to close the connection to the database
	StateClosing State = "closing"
	// StateClosed means the Connection object is disconnected to the database
	StateClosed State = "closed"
)

func (s State) String() string {
	return string(s)
}
