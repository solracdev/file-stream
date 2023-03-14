package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":8055")
	if err != nil {
		return err
	}

	_ = binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes over the network \n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		_ = sendFile(200000)
	}()

	server, err := NewFileServer("tcp", ":8055")
	if err != nil {
		log.Fatal(err)
	}
	server.start()
}
