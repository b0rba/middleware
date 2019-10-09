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
func (server Server) Run() {
	marshallerImpl := marshaller.Marshaller{}
	packetReply := packet.Packet{}
	var params []interface{}
	fmt.Println("Naming service on.")

	for {
		serverRequestHandlerImpl := srh.SRH{ServerHost: server.IP, ServerPort: server.Port}

		// Receive data
		receiveMsgBytes := (&serverRequestHandlerImpl).Receive()

		// 	unmarshall
		packetRequest := marshallerImpl.Unmarshall(receiveMsgBytes)

		// finding the operation
		operation := packetRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Lookup":
			packet := packetRequest.Bd.ReqBody.Body[0].(string)
			params = make([]interface{}, 2)
			params[0], params[1] = server.NS.Lookup(packet)
		case "Bind":
			packetBody := packetRequest.Bd.ReqBody.Body
			packet := packetBody[0].(string)
			packetBodyString := packetBody[1].(map[string]interface{})
			packet2 := clientproxy.InitClientProxy(packetBodyString["Host"].(string), int(packetBodyString["Port"].(float64)), int(packetBodyString["ID"].(float64)), packetBodyString["nameType"].(string))
			params = make([]interface{}, 1)
			params[0] = server.NS.Bind(packet, packet2)
			if params[0] != nil {
				params[0] = params[0].(error)
			}
		case "List":
			params = make([]interface{}, 1)
			params[0] = server.NS.List()
		}

		// assembly packetBody
		replyHeader := packet.ReplyHeader{Context: "", RequestID: packetRequest.Bd.ReqHeader.RequestID, Status: 1}
		replyBody := packet.ReplyBody{OperationResult: params}
		header := packet.Header{Magic: "packetBody", Version: "1.0", ByteOrder: true, MessageType: 0} // MessageType 0 = reply
		body := packet.Body{RepHeader: replyHeader, RepBody: replyBody}
		packetReply = packet.Packet{Hdr: header, Bd: body}

		// marshall reply
		msg2ClientBytes := marshallerImpl.Marshall(packetReply)

		// send Reply
		(&serverRequestHandlerImpl).Send(msg2ClientBytes)
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
	clientProxyMaps := make(map[string]clientproxy.ClientProxy)
	namingS := namingService.NamingService{Repository: clientProxyMaps}
	server := Server{NS: &namingS, IP: "localhost", Port: 8090}
	return server
}
