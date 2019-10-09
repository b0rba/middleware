package invoker

import (
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/server/distribution/lifecycle-management/pooling"
	"github.com/b0rba/middleware/my-middleware/server/infrastructure/srh"
	"github.com/b0rba/middleware/my-middleware/server/service/echoer"

	"fmt"
)

// EchoInvoker run the invoker.
type EchoInvoker struct {}

// Invoke starts the server
func (EchoInvoker) Invoke (){
	marshallerImpl := marshaller.Marshaller{}
	packetReply := packet.Packet{}
	replyParams := make([]interface{}, 1)

	// creating the pool
	echoList := make([]interface{}, 11)
	for i := 0; i < len(echoList); i++ {
		echoAux := echoer.Echo{}
		echoList[i] = &echoAux
	}
	echoPool := pooling.InitPool(echoList)
	defer pooling.EndPool(echoPool)

	fmt.Println("Server invoking.")

	for {
		serverRequestHandlerImpl := srh.SRH{ServerHost: "localhost", ServerPort:8080}
		
		// Receive data
		receiverMessageBytes := (&serverRequestHandlerImpl).Receive()

		// 	unmarshall
		packetPacketRequest := marshallerImpl.Unmarshall(receiverMessageBytes)

		// setup request
		packetRequest := int(packetPacketRequest.Bd.ReqBody.Body[0].(float64))
		var echo1 *echoer.Echo
		echo1 = echoPool.GetFromPool().(*echoer.Echo)

		// finding the operation
		operation := packetPacketRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Ech":
			replyParams[0] = echo1.Ech(packetRequest)
		}

		// assembly packet
		replyHeader := packet.ReplyHeader{Context: "", RequestID: packetPacketRequest.Bd.ReqHeader.RequestID, Status:1}
		replyBody := packet.ReplyBody{OperationResult: replyParams}
		header    := packet.Header{Magic:"packet", Version:"1.0", ByteOrder:true, MessageType:0} // MessageType 0 = reply
		body      := packet.Body{RepHeader: replyHeader, RepBody: replyBody}
		packetReply = packet.Packet{Hdr: header, Bd: body}

		// marshall reply
		msg2ClientBytes := marshallerImpl.Marshall(packetReply)

		// send Reply
		(&serverRequestHandlerImpl).Send(msg2ClientBytes)
	}
}