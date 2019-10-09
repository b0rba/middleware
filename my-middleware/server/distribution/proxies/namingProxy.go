package proxies

import (
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/common/service/namingService"
	"github.com/b0rba/middleware/my-middleware/server/infrastructure/srh"

	"fmt"
)

// Server managing a naming service.
type Server struct {
	NS   *namingService.NamingService
	IP string
	Port int
}

// Run runs the server.
func (server Server) Run() {
	marshallerImpl := marshaller.Marshaller{}
	packetReply := packet.Packet{}
	var params []interface{}
	fmt.Println("Naming service on.")

	for {
		serverRequestHandlerImpl := srh.SRH{ServerHost: server.IP, ServerPort: server.Port}

		// Receive data
		receiveMsgBytes := (&serverRequestHandlerImpl).Receive()

		// Unmarshall
		packetRequest := marshallerImpl.Unmarshall(receiveMsgBytes)

		// get the operation
		operation := packetRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Lookup":
			packetRqst := packetRequest.Bd.ReqBody.Body[0].(string)
			params = make([]interface{}, 2)
			params[0], params[1] = server.NS.Lookup(packetRqst) // get clientProxy from repo
		case "Bind":
			packetRqstBody := packetRequest.Bd.ReqBody.Body
			packetRqstBodyString := packetRqstBody[0].(string)
			packetBodyString := packetRqstBody[1].(map[string]interface{})
			packet2 := clientproxy.InitClientProxy(
				packetBodyString["Host"].(string),
				int(packetBodyString["Port"].(float64)),
				int(packetBodyString["ID"].(float64)),
				packetBodyString["nameType"].(string))
			params = make([]interface{}, 1)
			params[0] = server.NS.Bind(packetRqstBodyString, packet2)
			if params[0] != nil {
				params[0] = params[0].(error)
			}
		case "List":
			params = make([]interface{}, 1)
			params[0] = server.NS.List()
		}

		// assembly packetRqstBody
		replyHeader := packet.ReplyHeader{Context: "", RequestID: packetRequest.Bd.ReqHeader.RequestID, Status: 1}
		replyBody := packet.ReplyBody{OperationResult: params}
		header := packet.Header{Magic: "packetRqstBody", Version: "1.0", ByteOrder: true, MessageType: 0} // reply == 0
		body := packet.Body{RepHeader: replyHeader, RepBody: replyBody}
		packetReply = packet.Packet{Hdr: header, Bd: body}

		// marshall reply
		msg2ClientBytes := marshallerImpl.Marshall(packetReply)

		// send Reply
		(&serverRequestHandlerImpl).Send(msg2ClientBytes)
	}
}

// InitServer create the naming server.
func InitServer() Server {
	clientProxyMaps := make(map[string]clientproxy.ClientProxy)
	namingS := namingService.NamingService{Repository: clientProxyMaps}
	server := Server{NS: &namingS, IP: "localhost", Port: 8090}
	return server
}
