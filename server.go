package main

// Input is rpc input argument
type Input string

// Msisdn is the exported object
type Msisdn int

// Parse parses MSISDN input and returns MsisdnData
func (t *Msisdn) Parse(msisdn *Input, data *MsisdnData) error {
	m, err := ParseMsisdn(string(*msisdn))
	if err != nil {
		return err
	}

	*data = *m
	return nil
}
