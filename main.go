package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	for _, i := range []string{"38631123456", "38670987654", "8801812345678"} {
		err := parseMsisdn(i)
		if err != nil {
			log.Println(err)
		}
	}
}

var countryCodes = map[string]string{
	"386": "SI",
	"1":   "US",
	"880": "BD",
}

type ndcMno struct {
	ndc, mno string
}

var countryNDCs = map[string][]ndcMno{
	"SI": []ndcMno{
		ndcMno{"31", "Mobitel"},
		ndcMno{"41", "Mobitel"},
		ndcMno{"51", "Mobitel"},
		ndcMno{"71", "Mobitel"},
		ndcMno{"30", "Si.mobil"},
		ndcMno{"40", "Si.mobil"},
		ndcMno{"68", "Si.mobil"},
		ndcMno{"70", "Telemach"},
		ndcMno{"64", "T-2"},
	},
}

func parseMsisdn(in string) error {

	// assume that input is valid - just numbers
	if len(in) > 15 {
		return errors.New("Input too long.")
	}

	var isoID, cdc string
	for code, country := range countryCodes {
		if strings.HasPrefix(in, code) {
			cdc = code
			isoID = country
			in = strings.TrimPrefix(in, code)
			break
		}
	}

	if cdc == "" {
		return errors.New("Unknown Country Code.")
	}

	var mno string
	if countryData, ok := countryNDCs[isoID]; ok {
		for _, v := range countryData {
			if strings.HasPrefix(in, v.ndc) {
				mno = v.mno
				in = strings.TrimPrefix(in, v.ndc)
				break
			}
		}
	} else {
		return fmt.Errorf("Country \"%s\" not implemented", isoID)
	}

	if mno == "" {
		return errors.New("Network Destination Code unknown")
	}

	log.Println("cdc:", cdc, "country id:", isoID, "mno:", mno, "sn:", in)

	return nil
}
