package parser

import "testing"

func TestTrie(t *testing.T) {
	var root trie
	slo := &countryData{3, 2, "SI"}
	root.addCountry("386", slo)
	root.addCountry("385", &countryData{3, 2, "HR"})

	data, _ := root.findCountry("38631123123")
	if slo != data {
		t.Error("Trie find error")
	}

	root.addCountry("1", &countryData{1, 3, "US"})
	root.addCountry("1242", &countryData{4, 1, "BS"})

	data, _ = root.findCountry("12331123123")
	if data.isoID != "US" {
		t.Error("Trie find error - expected US, got ", data.isoID)
	}

	data, err := root.findCountry("712123")
	if err == nil {
		t.Error("No error returned for no match: ", data)
	}
}
