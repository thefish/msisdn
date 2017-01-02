package main

import (
	"errors"
	"log"
	"strings"
)

func main() {
	parse("38631123456")
	parse("8801812345678")
}

var countryCodes = map[string]string{
	"386": "SI",
	"1":   "US",
	"880": "BD",
}

func parse(msisdn string) error {

	// assume that input is valid - just numbers
	if len(msisdn) > 15 {
		return errors.New("Input too long.")
	}

	var isoID, cdc string
	for code, country := range countryCodes {
		if strings.HasPrefix(msisdn, code) {
			cdc = code
			isoID = country
			msisdn = strings.TrimLeft(msisdn, code)
			break
		}
	}

	log.Println("cdc:", cdc, "country id:", isoID, "ndc+sn:", msisdn)

	return nil
}
