package utils

// Invocation is a struct for invocation calls.
//
// Members:
//  Host    - the host target of the call.
//  Port    - the port of the call.
//  Request - the request.
//
type Invocation struct {
	Host    string
	Port    int
	Request Request
}

// Termination is a struct for terminate.
//
// Members:
//  Rep - the reply.
//
type Termination struct {
	Rep Reply
}

// IOR is a struct for holding a ID, a Port and the host.
//
// Members:
//  Host - the host.
//  Port - the port.
//  ID   - the id of the application.
//
type IOR struct {
	Host string
	Port int
	ID   int
}

// Request is a struct for the request data.
//
// Members:
//  Op     - the operation.
//  Params - list of params.
//
type Request struct {
	Op     string
	Params []interface{}
}

// Reply is a struct for the reply data.
//
// Members:
//  Result - the result of the operation operation.
//
type Reply struct {
	Result []interface{}
}
