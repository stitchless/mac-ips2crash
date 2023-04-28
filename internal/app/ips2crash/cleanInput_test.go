package ips2crash_test

import (
	"reflect"
	"testing"

	"github.com/jpeizer/mac-ips2crash/internal/app/ips2crash"
)

func TestIsJSONObject(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid JSON object",
			input:    `{"key": "value"}`,
			expected: true,
		},
		{
			name:     "Valid JSON object with nested object",
			input:    `{"key": {"key2": "value2"}}`,
			expected: true,
		},
		{
			name:     "Invalid JSON object",
			input:    `{"key": "value",}`,
			expected: false,
		},
		{
			name:     "Empty string",
			input:    ``,
			expected: false,
		},
		{
			name:     "JSON array",
			input:    `["value1", "value2"]`,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ips2crash.IsJSONObject(tc.input)
			if result != tc.expected {
				t.Errorf("IsJSONObject(%q) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestNewCrashReport(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ips2crash.InputFile
	}{
		{
			name:  "Valid input file with all sections",
			input: "Leading text\n{\"header\": \"value\"}\n{\n  \"payload\": {\n    \"nested\": \"data\"\n  }\n}\nTrailing text",
			expected: ips2crash.InputFile{
				LeadingText:  []string{"Leading text"},
				Header:       "{\"header\": \"value\"}",
				Payload:      "{  \"payload\": {    \"nested\": \"data\"  }}",
				TrailingText: []string{"Trailing text"},
			},
		},
		{
			name:  "Input file with only header",
			input: "{\"header\": \"value\"}",
			expected: ips2crash.InputFile{
				LeadingText:  nil,
				Header:       "{\"header\": \"value\"}",
				Payload:      "",
				TrailingText: nil,
			},
		},
		{
			name:  "Empty input file",
			input: "",
			expected: ips2crash.InputFile{
				LeadingText:  nil,
				Header:       "",
				Payload:      "",
				TrailingText: nil,
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ips2crash.NewCrashReport([]byte(tc.input))
			if err != nil {
				t.Errorf("NewCrashReport(%q) returned unexpected error: %v", tc.input, err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("NewCrashReport(%q) = %+v; want %+v", tc.input, result, tc.expected)
			}
		})
	}
}
