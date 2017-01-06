package parser

import (
	"errors"
	"fmt"
)

type trie struct {
	symbol   rune
	branches []*trie
	country  *countryData
}

type countryData struct {
	ccSize  int    // country code size (workaround for NANP specifics)
	mnoSize int    // number of digits in network identifier
	isoID   string // ISO 3166-1-alpha-2
}

func findBranch(r rune, branches []*trie) *trie {
	for _, b := range branches {
		if b.symbol == r {
			return b
		}
	}
	return nil
}

func (t *trie) addCountry(cc string, data *countryData) {
	var b *trie
	for _, r := range cc {
		b = findBranch(r, t.branches)
		if b == nil {
			b = new(trie)
			b.symbol = r
			t.branches = append(t.branches, b)
		}
		t = b
	}
	b.country = data
}

func (t *trie) printWhole(lvl int) {
	for i := 0; i < lvl; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%c", t.symbol)
	if t.country != nil {
		fmt.Printf(" %s %d ", t.country.isoID, t.country.mnoSize)
	}
	fmt.Println()
	for _, v := range t.branches {
		v.printWhole(lvl + 1)
	}
}

// findCountry finds a country in trie based on given string
func (t *trie) findCountry(in string) (country *countryData, e error) {
	var b, last *trie

	for _, r := range in {
		b = findBranch(r, t.branches)

		if b == nil {
			if t.country == nil {
				break
			}
			return t.country, nil
		}

		// remember last match, in case we traverse too deep
		if t.country != nil {
			last = t
		}
		t = b
	}

	if last != nil && last.country != nil {
		return last.country, nil
	}
	return nil, errors.New("Country code error.")
}
