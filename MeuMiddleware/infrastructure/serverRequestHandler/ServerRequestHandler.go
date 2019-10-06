package infrastructure

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

//ServerRequestHandler server2client setup struct
type ServerRequestHandler struct {
	ServerHost string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

//Send the []bytes
func (ServerRequestHandler) Send(msg2Client []byte) {
	//message size, message

	size := make([]byte, 4)
	l := uint32(len(msg2Client))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	_, err = conn.Write(msg2Client)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	conn.Close()
	ln.Close()
}
//Receive the []bytes
func (srh ServerRequestHandler) Receive() []byte {

	//listener, accept connections, message size, message
	ln, err = net.Listen("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	return msg
}
