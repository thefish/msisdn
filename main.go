package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	m, err := ParseMsisdn("38668123123")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(m)

	go initRPCServer()

	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer c.Close()

	client := jsonrpc.NewClient(c)

	var in Input
	var reply MsisdnData
	in = "3864123321"

	client.Call("Msisdn.Parse", &in, &reply)

	log.Println("rpc call: ", reply.String())
}

func initRPCServer() {
	ms := new(Msisdn)

	s := rpc.NewServer()
	err := s.Register(ms)
	if err != nil {
		log.Fatalf(err.Error())
	}

	s.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf(err.Error())
		}

		go s.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
