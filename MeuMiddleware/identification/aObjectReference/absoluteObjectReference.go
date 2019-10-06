package aObjectReference

// AbsoluteObjectReference references remote objects.
type AbsoluteObjectReference struct {
	IP        string //remote object's host IP
	Door      string //remote object's host door
	InvokerID int    //ID of the Invoker, used if there are more than one invoker.
	ObjectID  int    //ID of the remote object.
	Protocol  string //communication protocol.
}
