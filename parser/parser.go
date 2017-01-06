package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var trieRoot *trie

func init() {
	trieRoot = new(trie)

	for cc, data := range countries {
		trieRoot.addCountry(cc, data)
	}

	for _, v := range usaCC {
		trieRoot.addCountry(v, usaData)
	}

	for _, v := range canadaCC {
		trieRoot.addCountry(v, canadaData)
	}
}

// MsisdnData contains mobile network operator id, country dial code, subscriber
// number and country id
type MsisdnData struct {
	Mno, Cdc, Sn, CountryID string
}

func (m *MsisdnData) String() string {
	return fmt.Sprintf("mno:%s, cdc:%s, sn:%s, country id:%s", m.Mno, m.Cdc, m.Sn, m.CountryID)
}

// ParseMsisdn parses a MSISDN string and returns MsisdnData or an error.
func ParseMsisdn(in string) (*MsisdnData, error) {

	in = sanitize(in)

	// properly formatted msisdn has between 8 and 15 digits
	if ok, err := regexp.MatchString("^\\d{8,15}$", in); !ok || err != nil {
		return nil, errors.New("Invalid input")
	}

	cData, err := trieRoot.findCountry(in)
	if err != nil {
		return nil, err
	}
	offset := cData.ccSize
	snOffset := offset + cData.mnoSize

	cdc := in[:offset]
	mno := in[offset:snOffset]

	return &MsisdnData{mno, cdc, in[snOffset:], cData.isoID}, nil
}

// clean the input string
func sanitize(in string) string {

	// remove leading zeroes
	in = strings.TrimLeft(in, "0")

	// remove leading +
	in = strings.TrimPrefix(in, "+")

	// remove hyphens
	in = strings.Replace(in, "-", "", -1)

	// remove parens
	in = strings.Replace(in, "(", "", -1)
	in = strings.Replace(in, ")", "", -1)

	// remove spaces
	in = strings.Replace(in, " ", "", -1)
	return in
}
