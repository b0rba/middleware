package marshaller

import (
	"encoding/json"
	"log"

	"github.com/b0rba/middleware/my-middleware/common/distribution/packet"
)

// Marshaller enables Marshaller functions
type Marshaller struct{}

// Marshall is a funcion that receives a packet and transforms it to a bytes package
func (Marshaller) Marshall(message packet.Packet) []byte {
	resultBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}
	return resultBytes
}

// Unmarshall receives a byte package and transforms it to a packet
func (Marshaller) Unmarshall(message []byte) packet.Packet {
	resultBytes := packet.Packet{}
	err := json.Unmarshal(message, &resultBytes)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return resultBytes
}
