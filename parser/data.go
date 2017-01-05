package parser

var countryCodes = map[string]string{
	"386": "SI",
	"1":   "US",
	"880": "BD",
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
