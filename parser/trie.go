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
	ndcSize int
	isoID   string
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
		fmt.Printf(" %s %d ", t.country.isoID, t.country.ndcSize)
	}
	fmt.Println()
	for _, v := range t.branches {
		v.printWhole(lvl + 1)
	}
}

// findCountry finds a country in trie based on given string
func (t *trie) findCountry(in string) (country *countryData, cc int, e error) {
	var b, last *trie
	var lastInd int

	for i, r := range in {
		b = findBranch(r, t.branches)

		if b == nil {
			if t.country == nil {
				break
			}
			return t.country, i, nil
		}

		// remember last match, in case we traverse too deep
		if t.country != nil {
			last = t
			lastInd = i
		}
		t = b
	}

	if last != nil && last.country != nil {
		return last.country, lastInd, nil
	}
	return nil, 0, errors.New("Country code error.")
}
