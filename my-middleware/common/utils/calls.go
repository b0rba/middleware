package utils

// Invocation invocation calls.
type Invocation struct {
	Host    string
	Port    int
	Request Request
}

// Termination terminate.
type Termination struct {
	Rep Reply
}

// IOR holds a ID, a Port and the host.
type IOR struct {
	Host string
	Port int
	ID   int
}

// Request request data.
type Request struct {
	Op     string
	Params []interface{}
}

// Reply reply data.
type Reply struct {
	Result []interface{}
}
