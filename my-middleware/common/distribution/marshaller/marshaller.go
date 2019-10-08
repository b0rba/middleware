package marshaller

import (
	"encoding/json"
	"log"

	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
)

// Marshaller is a structure to enable Marshaller funcions
//
// Members:
//  none
//
type Marshaller struct{}

// Marshall is a funcion that receives a packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target packet
//
// Returns:
//  packet transformed to bytes
//
func (Marshaller) Marshall(msg packet.Packet) []byte {

	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

// Unmarshall is a funcion that receives a packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target bytes package
//
// Returns:
//   a packet
//
func (Marshaller) Unmarshall(msg []byte) packet.Packet {

	r := packet.Packet{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}
