package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

type FileServer struct {
	listener net.Listener
}

func NewFileServer(network, address string) (*FileServer, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return &FileServer{
		listener: ln,
	}, nil
}

func (s *FileServer) start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.readConn(conn)
	}
}

func (s *FileServer) readConn(conn net.Conn) {
	buff := bytes.NewBuffer(make([]byte, 0))
	for {
		var size int64
		_ = binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buff, conn, size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buff.Bytes())
		fmt.Printf("recieved %d bytes from connextion\n", n)
	}
}
