package ips2crash

import (
	"encoding/json"
	"strings"
)

type InputFile struct {
	LeadingText  []string
	Header       string
	Payload      string
	TrailingText []string
}

func isJSONObject(line string) bool {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(line), &js)
	return err == nil
}

func NewCrashReport(fileContent []byte) (InputFile, error) {
	var crashFile InputFile

	var nestedJSONInProgress bool
	bracesCount := 0
	headerFound := false

	lines := splitLines(fileContent)

	for _, line := range lines {
		if isJSONObject(line) {
			crashFile.Header = line
			headerFound = true
			continue
		}

		if !headerFound {
			crashFile.LeadingText = append(crashFile.LeadingText, line)
			continue
		}

		openBraces := strings.Count(line, "{")
		closeBraces := strings.Count(line, "}")

		if openBraces > 0 || closeBraces > 0 {
			nestedJSONInProgress = true
		}

		if nestedJSONInProgress {
			crashFile.Payload += line
			bracesCount += openBraces - closeBraces

			if bracesCount == 0 {
				nestedJSONInProgress = false
			}
		} else {
			crashFile.TrailingText = append(crashFile.TrailingText, line)
		}
	}

	return crashFile, nil
}
