package packet

// Packet is a structure for the default packet format.
//
// Members:
//  Hdr - the header of the packet.
//  Bd  - the body of the pacet.
//
type Packet struct {
	Hdr Header
	Bd  Body
}

// Header is a structure for holding the information of the header.
//
// Member:
//  Magic       - packet standard.
//  Version     - verion of the protocol.
//  ByteOrder   - byte ordering.
//  MessageType - the type of the message.
//  Size        - size of the packet.
//
type Header struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        int
}

// Body is a structure for holding the information of the body of the packet.
//
// Member:
//  ReqHeader   - body header for request.
//  ReqBody     - body body for request.
//  RepHeader   - body header for reply.
//  RepBody     - body body for reply.
//
type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

// RequestHeader is a structure for headers from requests.
//
// Member:
//  Context          - the context.
//  RequestID        - the id of the request.
//  ResponseExpected - flag for response.
//  ObjectKey        - the key of the object.
//  Operation        - the operation to be performed.
//
type RequestHeader struct {
	Context          string
	RequestID        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

// RequestBody is a structure for bodies from requests.
//
// Member:
//  Body - the body data.
//
type RequestBody struct {
	Body []interface{}
}

// ReplyHeader is a structure for headers from replies.
//
// Member:
//  Context   - the context.
//  RequestID - the id of the request.
//  Status    - the status.
//
type ReplyHeader struct {
	Context   string
	RequestID int
	Status    int
}

// ReplyBody is a structure for bodies from replies.
//
// Member:
//  OperationResult - the result of the request.
//
type ReplyBody struct {
	OperationResult []interface{}
}
