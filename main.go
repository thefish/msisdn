package main

import (
	"log"

	"msisdn/parser"
)

func main() {
	for _, i := range []string{"38631123456", "38670987654", "8801812345678"} {
		msisdn, err := parser.ParseMsisdn(i)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(msisdn)
	}
}
