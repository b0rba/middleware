package marshaller

import (
	"encoding/json"
	"log"
	"github.com/b0rba/middleware/MeuMiddleware/util/miop"
)

type Marshaller struct{}

func (Marshaller) Marshall(msg miop.Packet) []byte {

	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

func (Marshaller) Unmarshall(msg []byte) miop.Packet {

	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}