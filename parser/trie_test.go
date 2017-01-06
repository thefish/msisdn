package parser

import "testing"

func TestTrie(t *testing.T) {
	var root trie
	slo := &countryData{2, "SI"}
	root.addCountry("386", slo)
	root.addCountry("385", &countryData{2, "HR"})

	data, _, _ := root.findCountry("38631123123")
	if slo != data {
		t.Error("Trie find error")
	}

	root.addCountry("1", &countryData{3, "US"})
	root.addCountry("1242", &countryData{1, "BS"})

	data, i, _ := root.findCountry("12331123123")
	if data.isoID != "US" {
		t.Error("Trie find error - expected US, got ", data.isoID)
	}

	if i != 1 {
		t.Error("Trie wrong cc length, expected 1, got ", i)
	}

	data, i, err := root.findCountry("712123")
	if err == nil {
		t.Error("No error returned for no match: ", data, i)
	}
}
