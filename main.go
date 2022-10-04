package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"

	"msisdn/parser"
	"msisdn/server"
)

func main() {
	m, err := parser.ParseMsisdn("38668123123")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(m)

	go server.StartRPCServer(":8000")

	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer c.Close()

	client := jsonrpc.NewClient(c)

	var in server.Input
	var reply parser.MsisdnData
	in = "+386-41-23321"

	err = client.Call("Msisdn.Parse", &in, &reply)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("rpc call, in:%s, out:%s", in, reply.String())
}
