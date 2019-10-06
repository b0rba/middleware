package invoker

import (
	"github.com/b0rba/middleware/MeuMiddleware/utils/marshaller"
	"github.com/b0rba/middleware/MeuMiddleware/infrastructure/serverRequestHandler"
	"github.com/b0rba/middleware/MeuMiddleware/utils/miop"
	"github.com/b0rba/middleware/MeuMiddleware/utils/lifecycle_manager/pooling"
	"github.com/b0rba/middleware/MeuMiddleware/echo/impl"
)

// EchoInvoker is a structure to run the invoker.
//
// Members:
//  none
//
type EchoInvoker struct {}

// NewEchoInvoker is a funcion to initialize a EchoInvoker.
//
// Parameters:
//  none
//
// Returns:
//  the EchoInvoker
//
func NewEchoInvoker() EchoInvoker {
	p := new(EchoInvoker)
	return *p
}

// Invoke is a funcion to set the server running.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func (EchoInvoker) Invoke (){
	srhImpl := serverRequestHandler.ServerRequestHandler{ServerHost:"localhost",ServerPort:8080}
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	// creating the pool
	echoList := make([]interface{}, 11)
	for i := 0; i < len(echoList); i++ {
		echoList[i] = impl.Echo{}
	}
	echoPool := pooling.InitPool(echoList)
	defer pooling.EndPool(echoPool)

	for {
		// Receive data
		rcvMsgBytes := srhImpl.Receive()

		// 	unmarshall
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		
		
		// setup request
		_s1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
		var echo1 *impl.Echo
		echo1 = echoPool.GetFromPool().(*impl.Echo)

		// finding the operation
		operation := miopPacketRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Echo":
			replParams[0] = echo1.Echo(_s1)
		}

		// assembly packet
		repHeader := miop.ReplyHeader{Context:"", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status:1}
		repBody   := miop.ReplyBody{OperationResult: replParams}
		header    := miop.Header{Magic:"MIOP", Version:"1.0", ByteOrder:true, MessageType:1} // MessageType 1 = request
		body      := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// marshall reply
		msg2ClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send Reply
		srhImpl.Send(msg2ClientBytes)
	}
}