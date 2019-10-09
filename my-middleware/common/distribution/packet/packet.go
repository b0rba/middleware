package packet

// Packet default packet format.
type Packet struct {
	Hdr Header
	Bd  Body
}

// Header information of the header.
type Header struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        int
}

// Body of the packet.
type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

// RequestHeader headers from requests.
type RequestHeader struct {
	Context          string
	RequestID        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

// RequestBody bodies from requests.
type RequestBody struct {
	Body []interface{}
}

// ReplyHeader headers from replies.
type ReplyHeader struct {
	Context   string
	RequestID int
	Status    int
}

// ReplyBody bodies from replies.
type ReplyBody struct {
	OperationResult []interface{}
}
