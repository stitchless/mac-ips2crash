package ips2crash

func CleanOutput(crashReport IPSCrash) string {
	leadingText := crashReport.LeadingText
	payload := crashReport.FormattedReport
	trailingText := crashReport.TrailingText

	// combine the leading text []string, payload string, and trailing text []string
	// into a single string

	output := ""
	for _, line := range leadingText {
		output += line + "\n"
	}

	output += payload + "\n"

	for _, line := range trailingText {
		output += line + "\n"
	}
	
	return output
}
