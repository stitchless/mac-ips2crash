package ips2crash

import "testing"

func TestCleanOutput(t *testing.T) {
	testCases := []struct {
		name     string
		input    IPSCrash
		expected string
	}{
		{
			name: "Valid input file with all sections",
			input: IPSCrash{
				LeadingText: []string{
					"Leading text",
				},
				FormattedReport: "{\n  \"header\": \"value\"\n}",
				TrailingText: []string{
					"Trailing text",
				},
			},
			expected: "Leading text\n{\n  \"header\": \"value\"\n}\nTrailing text\n",
		},
		{
			name: "Input file with only header",
			input: IPSCrash{
				LeadingText:     nil,
				FormattedReport: "{\n  \"header\": \"value\"\n}",
				TrailingText:    nil,
			},
			expected: "{\n  \"header\": \"value\"\n}\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CleanOutput(tc.input)
			if result != tc.expected {
				t.Errorf("CleanOutput(%v) = %+v; want %+v", tc.input, result, tc.expected)
			}
		})
	}
}
