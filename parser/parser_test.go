package parser

import (
	"fmt"
	"testing"
)

func ExampleMsisdnPrintout() {
	m, _ := ParseMsisdn("38631123456")

	fmt.Print(m)
	// Output: mno:Mobitel, cdc:386, sn:123456, country id:SI
}

func TestParseMsisdn(t *testing.T) {
	m, _ := ParseMsisdn("38670987654")

	if m.Mno != "Telemach" {
		t.Error("MNO identifier invalid.", m.Mno)
	}

	if m.Cdc != "386" {
		t.Error("Invalid CDC.", m.Cdc)
	}

	if m.CountryID != "SI" {
		t.Error("Invalid country ID", m.CountryID)
	}

	if m.Sn != "987654" {
		t.Error("Invalid sn", m.Sn)
	}

}
