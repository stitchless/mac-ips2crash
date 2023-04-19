package ips2crash

import (
	"encoding/json"
)

func ParseCrashReport(crashReport InputFile) (IPSCrash, error) {
	var ipsCrash IPSCrash
	var err error

	ipsCrash.LeadingText = crashReport.LeadingText
	ipsCrash.TrailingText = crashReport.TrailingText

	ipsCrash.Header, err = parseHeader(crashReport.Header)
	if err != nil {
		return IPSCrash{}, err
	}

	ipsCrash.Payload, err = parsePayload(crashReport.Payload)
	if err != nil {
		return IPSCrash{}, err
	}

	return ipsCrash, nil
}

func parseHeader(header string) (IPSHeader, error) {
	var headerStruct IPSHeader

	err := json.Unmarshal([]byte(header), &headerStruct)
	if err != nil {
		return headerStruct, err
	}
	return headerStruct, nil
}

func parsePayload(payload string) (IPSPayload, error) {
	var crashReport IPSPayload
	err := json.Unmarshal([]byte(payload), &crashReport)
	if err != nil {
		return crashReport, err
	}
	return crashReport, nil
}
