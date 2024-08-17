package problem06

import ( 
	"encoding/binary"
	"fmt"
	"net"
)

type (
	TicketDispatcher struct {
		Roads []uint16 // Road IDs
		conn  net.Conn
	}
)

func (td *TicketDispatcher) UnmarshalBinary(data []byte) {
	offset := 1  // First byte is msgType header
	roadLen := 2 // uint16 2 bytes
	numRoads := data[offset]
	for i := 0; i < int(numRoads); i++ {
		pos := (i + offset) * roadLen
		road := binary.BigEndian.Uint16(data[pos : pos+roadLen])
		td.Roads = append(td.Roads, road)
	}
}

func (td *TicketDispatcher) send(t *Ticket) error {
	_, err := td.conn.Write(t.MarshalBinary())
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}