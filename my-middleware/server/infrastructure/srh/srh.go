package srh

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// SRH is a structure for Server to Client setups.
//
// Members:
//  ServerHost - server host IP.
//  ServerPort - server host port.
//  conn       - the connection of the srh.
//  ln         - the listener of the srh.
//
type SRH struct {
	ServerHost string
	ServerPort int
	conn       net.Conn
	ln         net.Listener
}

// Receive is a funcion that listens for messages
//
// Parameters:
// none
//
// Returns:
// Message received from client
//
func (srh *SRH) Receive() []byte {
	var err error

	// create listener
	srh.ln, err = net.Listen("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))
	utils.PrintError(err, "unable to listen on server request handler")

	// accept connections
	srh.conn, err = srh.ln.Accept()
	utils.PrintError(err, "unable to accept connection on server request handler")

	// receive message size
	size := make([]byte, 4)
	_, err = srh.conn.Read(size)
	utils.PrintError(err, "unable to read size on server request handler")

	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = srh.conn.Read(msg)
	utils.PrintError(err, "unable to read message on server request handler")

	return msg
}

// Send is a funcion that sends a message to a client
//
// Parameters:
// msgToClient - A byte package to be sent to a client
//
// Returns:
// none
//
func (srh *SRH) Send(msgToClient []byte) {
	// close connection
	defer srh.conn.Close()
	defer srh.ln.Close()

	// send message size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err := srh.conn.Write(size)
	utils.PrintError(err, "unable to write size to client on server request handler")

	// send message
	_, err = srh.conn.Write(msgToClient)
	utils.PrintError(err, "unable to write message to client on server request handler")
}
