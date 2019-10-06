package requestor

import (
	"github.com/b0rba/middleware/MeuMiddleware/infrastructure/clientRequestHandler"
	"github.com/b0rba/middleware/MeuMiddleware/marshaller"
	"github.com/b0rba/middleware/MeuMiddleware/util/aux"
	"github.com/b0rba/middleware/MeuMiddleware/util/miop"
)

// Requestor enable Requestor funcions
type Requestor struct{}

// Invoke receives a Invocation and returns a Interface based on the Invocation parameters
func (Requestor) Invoke(inv aux.Invocation) interface{} {
	marshallerInst := marshaller.Marshaller{}
	crhInst := clientRequestHandler.ClientRequestHandler{ServerHost: inv.Host, ServerPort: inv.Port}

	// create
	reqHeader := miop.RequestHeader{Context: "Context", RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 1} // 1 == Request
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	// serialise
	msg2ClientBytes := marshallerInst.Marshall(miopPacketRequest)

	// send request and receive reply
	msgFromServerBytes := crhInst.SendReceive(msg2ClientBytes)
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	// extract result
	r := miopPacketReply.Bd.RepBody.OperationResult

	return r
}
