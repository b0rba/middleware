package infrastructure

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

//ClientRequestHandler client2server setup struct
type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
}

//SendReceive sends and receives []bytes
func (crh ClientRequestHandler) SendReceive(msg2Server []byte) []byte {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", "localhost:"+strconv.Itoa(crh.ServerPort))
		if err == nil {
			break
		}
	}

	defer conn.Close()

	// send messsage size, message
	msg2ServerSize := make([]byte, 4)
	l := uint32(len(msg2Server))
	binary.LittleEndian.PutUint32(msg2ServerSize, l)
	conn.Write(msg2ServerSize)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	_, err = conn.Write(msg2Server)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	//receive message size, message
	msgFromServerSize := make([]byte, 4)
	_, err = conn.Read(msgFromServerSize)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(msgFromServerSize)

	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = conn.Read(msgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msgFromServer
}
