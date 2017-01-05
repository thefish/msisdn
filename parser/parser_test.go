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

func TestClean(t *testing.T) {
	if sanitize("0038631313131") != "38631313131" {
		t.Error("Should remove leading 0")
	}

	if sanitize("+38631313131") != "38631313131" {
		t.Error("Should remove leading +")
	}

	if sanitize("+386-313-13131") != "38631313131" {
		t.Error("Should remove hyphens")
	}

	if sanitize("38-(631)-313-131") != "38631313131" {
		t.Error("Should remove parens")
	}
}

func TestParseMsisdnExceptions(t *testing.T) {
	_, e := ParseMsisdn("1111111")
	if e == nil {
		t.Error("Shouldn't allow number with less than 8 digits")
	}

	_, e = ParseMsisdn("1111111111111111")
	if e == nil {
		t.Error("Shouldn't allow number with more than 15 digits")
	}

	_, e = ParseMsisdn("386a1123123")
	if e == nil {
		t.Error("Non-digits in number.")
	}

}
