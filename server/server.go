package server

import (
	"msisdn/parser"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Input is rpc input argument
type Input string

// Msisdn is the exported object
type Msisdn int

// Parse parses MSISDN input and returns MsisdnData
func (t *Msisdn) Parse(msisdn *Input, data *parser.MsisdnData) error {
	m, err := parser.ParseMsisdn(string(*msisdn))
	if err != nil {
		return err
	}

	*data = *m
	return nil
}

// StartRPCServer starts the rpc server on given port number
func StartRPCServer(port string) error {
	ms := new(Msisdn)

	s := rpc.NewServer()
	err := s.Register(ms)
	if err != nil {
		return err
	}

	s.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go s.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
