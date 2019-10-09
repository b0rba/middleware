package srh

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// SRH Server to Client setups.
type SRH struct {
	ServerHost string
	ServerPort int
	conn       net.Conn
	listener   net.Listener
}

// Receive listener for messages
func (serverRequestHandler *SRH) Receive() []byte {
	var err error

	// create listener
	serverRequestHandler.listener, err =
		net.Listen("tcp", serverRequestHandler.ServerHost+":"+strconv.Itoa(serverRequestHandler.ServerPort))
	utils.PrintError(err, "unable to listen on SRH")

	// accept connections
	serverRequestHandler.conn, err = serverRequestHandler.listener.Accept()
	utils.PrintError(err, "unable to accept connection on SRH")

	// receive message size
	size := make([]byte, 4)
	_, err = serverRequestHandler.conn.Read(size)
	utils.PrintError(err, "unable to read message size on SRH")

	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	message := make([]byte, sizeInt)
	_, err = serverRequestHandler.conn.Read(message)
	utils.PrintError(err, "unable to read message on SRH")

	return message
}

// Send sends a message to a client
func (serverRequestHandler *SRH) Send(msgToClient []byte) {
	defer serverRequestHandler.conn.Close()
	defer serverRequestHandler.listener.Close()

	// send message size
	size := make([]byte, 4)
	length := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, length)
	_, err := serverRequestHandler.conn.Write(size)
	utils.PrintError(err, "unable to send message size on SRH")

	// send message
	_, err = serverRequestHandler.conn.Write(msgToClient)
	utils.PrintError(err, "unable to send message on SRH")
}
