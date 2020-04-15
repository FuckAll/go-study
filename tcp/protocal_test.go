package tcp

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestPackage(t *testing.T) {
	bytes := Packet(1, 1, 1, "你好啊")
	fmt.Println(string(bytes))
}

func TestTcpListen(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		t.Error(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Error(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			packet, err2 := UnPacket(c)
			if err2 != nil {
				fmt.Println("unpacket error:", err2.Error())
				return
			}
			fmt.Println("获取数据:", packet)
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁\n"))
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁"))
			_, err2 = c.Write([]byte("你好么老铁"))
			if err2 != nil {
				fmt.Println("write error:", err2.Error())
			}
		}(conn)

	}
}

func TestTcpDial(t *testing.T) {
	dial, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		t.Error(err)
	}

	dial.Write(Packet(1, 1, 1, "你可以啊老铁"))
	readString, _ := bufio.NewReader(dial).ReadString('\n')

	//readString, _ := bufio.NewReader(dial).ReadString('\n')
	fmt.Printf("收到信息:%s\n", readString)
	dial.Close()
}
