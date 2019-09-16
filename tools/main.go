package main

import (
	"bytes"
	"fmt"

	"github.com/flyaways/header-quic/protocol"
	"github.com/flyaways/header-quic/wire"
)

func main() {
	//func handlePacket(data []byte, shortHeaderConnIDLen int, sentBy protocol.Perspective, version protocol.VersionNumber) (*protocol.Header, error) {
	// Header is the header of a QUIC packet.
	// It contains fields that are only needed for the gQUIC Public Header and the IETF draft Header.
	// type Header struct {
	// 	IsPublicHeader bool

	// 	Raw []byte

	// 	Version protocol.VersionNumber

	// 	DestConnectionID     protocol.ConnectionID
	// 	SrcConnectionID      protocol.ConnectionID
	// 	OrigDestConnectionID protocol.ConnectionID // only needed in the Retry packet

	// 	PacketNumberLen protocol.PacketNumberLen
	// 	PacketNumber    protocol.PacketNumber

	// 	IsVersionNegotiation bool
	// 	SupportedVersions    []protocol.VersionNumber // Version Number sent in a Version Negotiation Packet by the server

	// 	// only needed for the gQUIC Public Header
	// 	VersionFlag          bool
	// 	ResetFlag            bool
	// 	DiversificationNonce []byte

	// 	// only needed for the IETF Header
	// 	Type         protocol.PacketType
	// 	IsLongHeader bool
	// 	KeyPhase     int
	// 	PayloadLen   protocol.ByteCount
	// 	Token        []byte
	// }
	header, err := handlePacket([]byte(""), 4, protocol.PerspectiveClient, protocol.Version39)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(header)

}

func handlePacket(data []byte, shortHeaderConnIDLen int, sentBy protocol.Perspective, version protocol.VersionNumber) (*wire.Header, error) {
	r := bytes.NewReader(data)
	iHdr, err := wire.ParseInvariantHeader(r, shortHeaderConnIDLen)
	// drop the packet if we can't parse the header
	if err != nil {
		return nil, fmt.Errorf("error parsing invariant header: %s", err)
	}

	hdr, err := iHdr.Parse(r, sentBy, version)
	if err != nil {
		return nil, fmt.Errorf("error parsing header: %s", err)
	}
	hdr.Raw = data[:len(data)-r.Len()]
	packetData := data[len(data)-r.Len():]

	if hdr.IsLongHeader && hdr.Version.UsesLengthInHeader() {
		if protocol.ByteCount(len(packetData)) < hdr.PayloadLen {
			return nil, fmt.Errorf("packet payload (%d bytes) is smaller than the expected payload length (%d bytes)", len(packetData), hdr.PayloadLen)
		}
	}

	return hdr, nil
}
