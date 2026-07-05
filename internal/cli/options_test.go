package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseOptions(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		assert func(t *testing.T, options options)
	}{
		{
			name: "practice number",
			args: []string{"3"},
			assert: func(t *testing.T, options options) {
				if !options.hasPracticeNumber || options.practiceNumber != 3 {
					t.Fatalf("got %+v, want practice number 3", options)
				}
			},
		},
		{
			name: "list",
			args: []string{"--list"},
			assert: func(t *testing.T, options options) {
				if !options.list {
					t.Fatalf("got %+v, want list enabled", options)
				}
			},
		},
		{
			name: "resume",
			args: []string{"--resume"},
			assert: func(t *testing.T, options options) {
				if !options.resume {
					t.Fatalf("got %+v, want resume enabled", options)
				}
			},
		},
		{
			name: "reset progress",
			args: []string{"--reset-progress"},
			assert: func(t *testing.T, options options) {
				if !options.resetProgress {
					t.Fatalf("got %+v, want reset progress enabled", options)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := parseOptions(test.args, &bytes.Buffer{})
			if err != nil {
				t.Fatal(err)
			}
			test.assert(t, parsed)
		})
	}
}

func TestParseOptionsRejectsInvalidCombinations(t *testing.T) {
	tests := [][]string{
		{"not-a-number"},
		{"1", "2"},
		{"--list", "--resume"},
		{"--resume", "3"},
	}

	for _, args := range tests {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			if _, err := parseOptions(args, &bytes.Buffer{}); err == nil {
				t.Fatalf("parseOptions(%q) returned no error", args)
			}
		})
	}
}
