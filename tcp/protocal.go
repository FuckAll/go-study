package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/*
	协议格式：20字节长度为header
	----------------------------------------------------------------
	headerLength | clientVersion | cmdId | seq | bodyLength | body|
*/

const HEADER_LENGTH = 20

func Packet(clientVersion uint32, cmdId uint32, seq uint32, body string) []byte {
	buffer := make([]byte, HEADER_LENGTH+len(body))
	// headLength
	binary.BigEndian.PutUint32(buffer[0:4], HEADER_LENGTH)
	// clientVersion
	binary.BigEndian.PutUint32(buffer[4:8], clientVersion)
	// cmdId
	binary.BigEndian.PutUint32(buffer[8:12], cmdId)
	// seq
	binary.BigEndian.PutUint32(buffer[12:16], seq)
	// bodyLength
	binary.BigEndian.PutUint32(buffer[16:20], uint32(len(body)))
	// body
	copy(buffer[20:], body)
	fmt.Println(len(buffer))
	return buffer
}

type Protocol struct {
	headLength    uint32
	clientVersion uint32
	cmdId         uint32
	seq           uint32
	bodyLength    uint32
	body          []byte
}

func UnPacket(c net.Conn) (*Protocol, error) {
	var p = &Protocol{}
	header := make([]byte, HEADER_LENGTH)
	_, err := io.ReadFull(c, header)
	if err != nil {
		return p, err
	}
	p.headLength = binary.BigEndian.Uint32(header[0:4])
	p.clientVersion = binary.BigEndian.Uint32(header[4:8])
	p.cmdId = binary.BigEndian.Uint32(header[8:12])
	p.seq = binary.BigEndian.Uint32(header[12:16])
	p.bodyLength = binary.BigEndian.Uint32(header[16:20])
	body := make([]byte, p.bodyLength)
	_, err = io.ReadFull(c, body)
	if err != nil {
		return p, err
	}
	p.body = body
	return p, nil
}
