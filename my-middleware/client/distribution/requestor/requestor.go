package requestor

import (
	"github.com/b0rba/middleware/my-middleware/client/infrastructure/crh"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Requestor is a structure to enable Requestor funcions
//
// Members:
//  none
//
type Requestor struct{}

// Invoke is a funcion that receives a Invocation and returns a Interface based on the Invocation parameters
//
// Parameters:
// inv - Received invocation
//
// Returns:
// interface
//
func (Requestor) Invoke(inv utils.Invocation) interface{} {
	marshallerInst := marshaller.Marshaller{}
	crhInst := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	// create request packet
	reqHeader := packet.RequestHeader{
		Context:"Context",
		RequestID: 1000,
		ResponseExpected: true,
		ObjectKey: 2000,
		Operation: inv.Request.Op}
	reqBody := packet.RequestBody{Body: inv.Request.Params}
	header := packet.Header{Magic: "packet", Version: "1.0", ByteOrder:true, MessageType:1 } // MessageType = 1 == Request
	body := packet.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	packetPacketRequest := packet.Packet{Hdr: header, Bd: body}

	// serialise request packet
	msgToClientBytes := marshallerInst.Marshall(packetPacketRequest)
	
	// send request packet and receive reply packet
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes)

	packetPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)
	
	// extract result from reply packet
	r := packetPacketReply.Bd.RepBody.OperationResult

	return r
}