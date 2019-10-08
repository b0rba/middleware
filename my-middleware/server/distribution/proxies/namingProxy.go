package proxies

import (
	"github.com/b0rba/middleware/my-middleware/common/service/namingService"
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/server/infrastructure/srh"

	"fmt"
)

// Server is a structure for managing a naming service.
//
// Members:
//  NS   - the naming service.
//  IP   - the ip of the server.
//  Port - port to the service.
//
type Server struct {
	NS   *namingService.NamingService
	IP string
	Port int
}

// Run is a function to run the server.
//
// parameters:
//  none.
//
// Returns:
//  none
//
func (sv Server) Run() {
	marshallerImpl := marshaller.Marshaller{}
	packetPacketReply := packet.Packet{}
	var replParams []interface{}
	fmt.Println("Naming service on.")

	for {
		srhImpl := srh.SRH{ServerHost: sv.IP, ServerPort: sv.Port}

		// Receive data
		rcvMsgBytes := (&srhImpl).Receive()

		// 	unmarshall
		packetPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)

		// finding the operation
		operation := packetPacketRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Lookup":
			p1 := packetPacketRequest.Bd.ReqBody.Body[0].(string)
			replParams = make([]interface{}, 2)
			replParams[0], replParams[1] = sv.NS.Lookup(p1)
		case "Bind":
			bd := packetPacketRequest.Bd.ReqBody.Body
			p1 := bd[0].(string)
			bdConv := bd[1].(map[string]interface{})
			p2 := clientproxy.InitClientProxy(bdConv["Host"].(string), int(bdConv["Port"].(float64)), int(bdConv["ID"].(float64)), bdConv["TypeName"].(string))
			replParams = make([]interface{}, 1)
			replParams[0] = sv.NS.Bind(p1, p2)
			if replParams[0] != nil {
				replParams[0] = replParams[0].(error)
			}
		case "List":
			replParams = make([]interface{}, 1)
			replParams[0] = sv.NS.List()
		}

		// assembly packet
		repHeader := packet.ReplyHeader{Context: "", RequestID: packetPacketRequest.Bd.ReqHeader.RequestID, Status: 1}
		repBody := packet.ReplyBody{OperationResult: replParams}
		header := packet.Header{Magic: "packet", Version: "1.0", ByteOrder: true, MessageType: 0} // MessageType 0 = reply
		body := packet.Body{RepHeader: repHeader, RepBody: repBody}
		packetPacketReply = packet.Packet{Hdr: header, Bd: body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(packetPacketReply)

		// send Reply
		(&srhImpl).Send(msgToClientBytes)
	}
}

// InitServer is a function to create the naming server.
//
// parameters:
//  none.
//
// Returns:
//  the running server.
//
func InitServer() Server {
	cpMap := make(map[string]clientproxy.ClientProxy)
	ns := namingService.NamingService{Repository: cpMap}
	sv := Server{NS: &ns, IP: "localhost", Port: 8090}
	return sv
}
