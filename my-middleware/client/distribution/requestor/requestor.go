package requestor

import (
	"github.com/b0rba/middleware/my-middleware/client/infrastructure/crh"
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Requestor enables Requestor funcions
type Requestor struct{}

// Invoke receives a Invocation and returns a Interface based on the Invocation parameters
func (Requestor) Invoke(inv utils.Invocation) interface{} {
	marshallerInst := marshaller.Marshaller{}
	clientRequestHandlerInst := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	// create request packet
	requestHeader := packet.RequestHeader{
		Context:"Context",
		RequestID: 1000,
		ResponseExpected: true,
		ObjectKey: 2000,
		Operation: inv.Request.Op}
	requestBody := packet.RequestBody{Body: inv.Request.Params}
	header := packet.Header{Magic: "packet", Version: "1.0", ByteOrder:true, MessageType:1 } // MessageType = 1 == Request
	body := packet.Body{ReqHeader: requestHeader, ReqBody: requestBody}
	packetPacketRequest := packet.Packet{Hdr: header, Bd: body}

	// serialise request packet
	msg2ClientBytes := marshallerInst.Marshall(packetPacketRequest)
	
	// send request packet and receive reply packet
	msgFromServerBytes := clientRequestHandlerInst.SendReceive(msg2ClientBytes)

	packetPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)
	
	// extract result from reply packet
	result := packetPacketReply.Bd.RepBody.OperationResult

	return result
}