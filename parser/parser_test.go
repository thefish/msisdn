package parser

import (
	"fmt"
	"testing"
)

func ExampleMsisdnPrintout() {
	m, _ := ParseMsisdn("38631123456")

	fmt.Print(m)
	// Output: mno:31, cdc:386, sn:123456, country id:SI
}

func TestParseMsisdn(t *testing.T) {
	m, _ := ParseMsisdn("38670987654")

	if m.Mno != "70" {
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

func TestParseMsisdnNANP(t *testing.T) {
	// USA
	m, _ := ParseMsisdn("1801-433-7300")
	if m.CountryID != "US" {
		t.Error("Country identifier invalid. Expected US, got", m.CountryID)
	}

	if m.Cdc != "1" {
		t.Error("Invalid CDC. Expected 1, got", m.Cdc)
	}

	if m.Sn != "4337300" {
		t.Error("Invalid sn. Expected 4337300, got", m.Sn)
	}

	// Canada
	m, _ = ParseMsisdn("1604-522-6600")
	if m.CountryID != "CA" {
		t.Error("MNO identifier invalid. Expected CA, got", m.CountryID)
	}

	if m.Cdc != "1" {
		t.Error("Invalid CDC. Expected 1, got", m.Cdc)
	}

	if m.Sn != "5226600" {
		t.Error("Invalid sn. Expected 5226600, got", m.Sn)
	}

	if m.Mno != "604" {
		t.Error("Invalid mno. Expected 604, got", m.Sn)
	}

	// another country from NANP - The Bahamas
	m, _ = ParseMsisdn("+1 242 123123")
	if m.CountryID != "BS" {
		t.Error("MNO identifier invalid. Expected BS, got", m.CountryID)
	}

	if m.Cdc != "1242" {
		t.Error("Invalid CDC. Expected 1242, got", m.Cdc)
	}

	if m.Sn != "123123" {
		t.Error("Invalid sn. Expected 123123, got", m.Sn)
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

	if sanitize("38 (631) 313 131") != "38631313131" {
		t.Error("Should remove spaces")
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
