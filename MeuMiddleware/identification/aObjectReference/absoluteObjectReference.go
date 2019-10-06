package aObjectReference

// AbsoluteObjectReference is a structure for referencing remote objects.
//
// Members:
//  IP        - remote object's host IP.
//  Door      - remote object's host door.
//  InvokerID - ID of the Invoker, used if there are more than one invoker.
//  ObjectID  - ID of the remote object.
//  Protocol  - communication protocol.
//
type AbsoluteObjectReference struct {
	IP        string
	Door      string
	InvokerID int
	ObjectID  int
	Protocol  string
}
