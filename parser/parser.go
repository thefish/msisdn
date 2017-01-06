package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// MsisdnData contains mobile network operator id, country dial code, subscriber
// number and country id
type MsisdnData struct {
	Mno, Cdc, Sn, CountryID string
}

func (m *MsisdnData) String() string {
	return fmt.Sprintf("mno:%s, cdc:%s, sn:%s, country id:%s", m.Mno, m.Cdc, m.Sn, m.CountryID)
}

// ParseMsisdn parses a MSISDN string and returns MsisdnData or an error.
func (t *trie) ParseMsisdn(in string) (*MsisdnData, error) {

	in = sanitize(in)

	// properly formatted msisdn has between 8 and 15 digits
	if ok, err := regexp.MatchString("^\\d{8,15}$", in); !ok || err != nil {
		return nil, errors.New("Invalid input")
	}

	cData, err := t.findCountry(in)
	if err != nil {
		return nil, err
	}
	offset := cData.ccSize

	cdc := in[:offset]
	snOffset := offset + cData.mnoSize
	mno := in[offset:snOffset]

	return &MsisdnData{mno, cdc, in[snOffset:], cData.isoID}, nil
	// if strings.HasPrefix(in, "1") {
	// 	return parseNanpMsisdn(in)
	// }

	// var isoID, cdc string
	// for code, country := range countryCodes {
	// 	if strings.HasPrefix(in, code) {
	// 		cdc = code
	// 		isoID = country
	// 		in = strings.TrimPrefix(in, code)
	// 		break
	// 	}
	// }

	// if cdc == "" {
	// 	return nil, errors.New("Unknown Country Code.")
	// }

	// var mno string
	// if countryData, ok := countryNDCs[isoID]; ok {
	// 	for _, v := range countryData {
	// 		if strings.HasPrefix(in, v.ndc) {
	// 			mno = v.mno
	// 			in = strings.TrimPrefix(in, v.ndc)
	// 			break
	// 		}
	// 	}
	// } else {
	// 	return nil, fmt.Errorf("Country \"%s\" not implemented", isoID)
	// }

	// if mno == "" {
	// 	return nil, errors.New("Network Destination Code unknown")
	// }

	// return &MsisdnData{mno, cdc, in, isoID}, nil
}

const usaRgx = "^1(201|202|203|205|206|207|208|209|210|212|213|214|215|216|217|218|219|224|225|228|229|231|234|239|240|248|251|252|253|254|256|260|262|267|269|270|272|276|281|301|302|303|304|305|307308|309|310|312|313|314|315|316|317|318|319|320|321|321|323|325|330|331|334|336|337|339|346|347|351|352|360|361|364|385|386|401|402|404|405|406|407|408|409|410|412|413|414|415|417|419|423|424|425|430|432|434|435|440|442|443|464|469|470|470|475|478|479|480|484|501|502|503|504|505|507|508|509|510|512|513|515|516|517|518|520|530|531|539|540|541|551|559|561|562|563|567|570|571|573|574|575|580|585|586|601|602|603|605|606|607|608|609|610|612|614|615|616|617|618|619|620|623|626|630|631|636|641|646|650|651|657|660|661|662|669|678|678|681|682|701|702|703|704|706|707|708|710|712|713|714|715|716|717|718|719|720|724|725|727|731|732|734|740|747|754|757|760|762|763|765|769|770|772|773|774|775|779|781|785|786|801|802|803|804|805|808|810|812|813|814|815|816|817|818|828|830|831|832|843|845|847|848|850|856|857|858|860|862|863|864|865|869|870|878|878|901|903|904|906|907|908|909|910|912|913|914|915|916|917|917|918|919|920|925|928|930|931|936|937|938|940|941|947|949|951|952|954|956|970|971|972|973|978|979|980|984|985|989)"

const canadaRgx = "^1(204|226|236|236|249|250|289|306|343|365|403|416|418|431|437|438|450|506|514|519|579|581|587|587|604|613|639|647|705|709|778|778|780|782|782|807|819|873|902|902|905)"

// parse MSISDN number according to North American Numbering Plan
func parseNanpMsisdn(in string) (*MsisdnData, error) {

	if m, _ := regexp.MatchString(usaRgx, in); m {
		out := in[1:]
		return &MsisdnData{"", "1", out, "US"}, nil
	}

	if m, _ := regexp.MatchString(canadaRgx, in); m {
		out := in[1:]
		return &MsisdnData{"", "1", out, "CA"}, nil
	}

	var isoID, cdc string
	for code, country := range nanpCountryCodes {
		if strings.HasPrefix(in, code) {
			cdc = code
			isoID = country
			in = strings.TrimPrefix(in, code)
			break
		}
	}

	if cdc == "" {
		return nil, errors.New("Unknown Country Code.")
	}

	return &MsisdnData{"", cdc, in, isoID}, nil
}

// sanitize cleans the input
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
