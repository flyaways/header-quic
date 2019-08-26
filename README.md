# header-quic

## Separated parsePacket from https://github.com/lucas-clemente/quic-go 

> for quic learn and test,and maybe modify it later.

> Thanks a million to https://github.com/lucas-clemente/quic-go

```go
package main

import (
	"fmt"

	"github.com/flyaways/header-quic/wire"
)

func main() {
	// ParsePacket parses a packet.
	// If the packet has a long header, the packet is cut according to the length field.
	// If we understand the version, the packet is header up unto the packet number.
	// Otherwise, only the invariant part of the header is parsed.
	//func ParsePacket(data []byte, shortHeaderConnIDLen int) (*Header, []byte /* packet data */, []byte /* rest */, error)
	fmt.Println(wire.ParsePacket([]byte(""), 4))
}
```
