package parser

var countryCodes = map[string]string{
	"386": "SI",
	"880": "BD",
}

var nanpCountryCodes = map[string]string{
	"1242": "BS",
	"1246": "BB",
	"1264": "AI",
	"1268": "AG",
	"1340": "VI",
	"1345": "KY",
	"1441": "BM",
	"1473": "GD",
	"1664": "MS",
	"1671": "GU",
	"1758": "LC",
	"1767": "DM",
	"1784": "VC",
	"1809": "DO",
	"1868": "TT",
	"1869": "KN",
	"1876": "JM",
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
	"BD": []ndcMno{
		ndcMno{"11", "Citycell"},
		ndcMno{"13", "Grameenphone"},
		ndcMno{"15", "TeleTalk"},
		ndcMno{"16", "Robi"},
		ndcMno{"17", "Grameenphone"},
		ndcMno{"18", "Robi"},
		ndcMno{"19", "Banglalink"},
	},
}

var countries = map[string]*countryData{
	"386": &countryData{3, 2, "SI"},
	"880": &countryData{3, 2, "BD"},
	"44":  &countryData{2, 3, "GB"},
	"1":   &countryData{1, 3, "US"},
}

func initCountryData() *trie {
	t := new(trie)

	for cc, data := range countries {
		t.addCountry(cc, data)
	}

	return t
}
