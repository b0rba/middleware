package invoker

import (
	"github.com/b0rba/middleware/my-middleware/common/distribution/marshaller"
	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
	"github.com/b0rba/middleware/my-middleware/server/infrastructure/srh"
	"github.com/b0rba/middleware/my-middleware/server/distribution/lifecycle-management/pooling"
	"github.com/b0rba/middleware/my-middleware/server/service/calculator"

	"fmt"
)

// CalculatorInvoker is a structure to run the invoker.
//
// Members:
//  none
//
type CalculatorInvoker struct {}

// Invoke is a funcion to set the server running.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func (CalculatorInvoker) Invoke (){
	marshallerImpl := marshaller.Marshaller{}
	packetPacketReply := packet.Packet{}
	replParams := make([]interface{}, 1)

	// creating the pool
	calculatorList := make([]interface{}, 11)
	for i := 0; i < len(calculatorList); i++ {
		calcAux := calculator.Calculator{}
		calculatorList[i] = &calcAux
	}
	multPool := pooling.InitPool(calculatorList)
	defer pooling.EndPool(multPool)

	fmt.Println("Server invoking.")

	for {
		srhImpl := srh.SRH{ServerHost:"localhost", ServerPort:8080}
		
		// Receive data
		rcvMsgBytes := (&srhImpl).Receive()

		// 	unmarshall
		packetPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		
		
		// setup request
		_p1 := int(packetPacketRequest.Bd.ReqBody.Body[0].(float64))
		var multA *calculator.Calculator
		multA = multPool.GetFromPool().(*calculator.Calculator)

		// finding the operation
		operation := packetPacketRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Mul":
			replParams[0] = multA.Mul(_p1)
		}

		// assembly packet
		repHeader := packet.ReplyHeader{Context:"", RequestID: packetPacketRequest.Bd.ReqHeader.RequestID, Status:1}
		repBody   := packet.ReplyBody{OperationResult: replParams}
		header    := packet.Header{Magic:"packet", Version:"1.0", ByteOrder:true, MessageType:0} // MessageType 0 = reply
		body      := packet.Body{RepHeader: repHeader, RepBody: repBody}
		packetPacketReply = packet.Packet{Hdr: header, Bd: body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(packetPacketReply)

		// send Reply
		(&srhImpl).Send(msgToClientBytes)
	}
}