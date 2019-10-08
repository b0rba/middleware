package crh

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// CRH is a structure for Client to Server setups.
//
// Members:
//  ServerHost - server host IP.
//  ServerPort - server host port.
//
type CRH struct {
	ServerHost string
	ServerPort int
}

// SendReceive is a funcion that receives a byte package and sends it to a server
//
// Parameters:
//  msgToServer - Package to be sent
//
// Returns:
//  Message received from server
//
func (crh CRH) SendReceive(msgToServer []byte) []byte {

	// connect to server
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", "localhost:"+strconv.Itoa(crh.ServerPort))
		if err == nil {
			// connected to server.
			break
		}

	}
	defer conn.Close()

	// send message size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToServer))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	conn.Write(sizeMsgToServer)
	utils.PrintError(err, "unable to write size to server on client request handler")

	// send message
	_, err = conn.Write(msgToServer)
	utils.PrintError(err, "unable to write message to server on client request handler")

	// receive message size
	sizeMsgFromServer := make([]byte, 4)
	_, err = conn.Read(sizeMsgFromServer)
	utils.PrintError(err, "unable to read size from server on client request handler")

	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	// receive reply
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = conn.Read(msgFromServer)
	utils.PrintError(err, "unable to read message from server on client request handler")

	return msgFromServer
}
