package parser

import (
	"errors"
	"fmt"
	"strings"
)

// MsisdnData contains mobile network operator id, country dial code, subscriber
// number and country id
type MsisdnData struct {
	Mno, Cdc, Sn, CountryID string
}

func (m *MsisdnData) String() string {
	return fmt.Sprintf("mno:%s, cdc:%s, sn:%s, country id:%s", m.Mno, m.Cdc, m.Sn, m.CountryID)
}

// ParseMsisdn parses a valid MSISDN string and returns MsisdnData or an error.
func ParseMsisdn(in string) (*MsisdnData, error) {

	// assume that input is valid - just numbers
	if len(in) > 15 {
		return nil, errors.New("Input too long.")
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
		return nil, errors.New("Unknown Country Code.")
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
		return nil, fmt.Errorf("Country \"%s\" not implemented", isoID)
	}

	if mno == "" {
		return nil, errors.New("Network Destination Code unknown")
	}

	return &MsisdnData{mno, cdc, in, isoID}, nil
}
