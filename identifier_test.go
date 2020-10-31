package deckard

import (
	"encoding/json"
	"testing"
	//"github.com/jasimmons/deckard"
)

func TestMarshal(t *testing.T) {
	tt := []struct {
		name      string
		in        *Identifier
		expected  string
		expectErr bool
	}{
		{
			name: "marshal valid Identifier",
			in: &Identifier{
				ID:   1,
				Name: "example identifier",
				Tags: map[string]struct{}{
					"foo": struct{}{},
					"bar": struct{}{},
				},
				Endpoint: "http://example.com:80",
			},
			expected: `{
"id": 1,
"name": "example identifier",
"tags": [
"foo",
"bar"
],
"endpoint": "http://example.com:80"
}`,
			expectErr: false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			_, err := json.Marshal(test.in)
			if err != nil && !test.expectErr {
				t.Errorf("expected no error\ngot error %v", err)
			}
			// TODO: compare test.expected vs. actual
		})
	}
}

func TestParseSchedule(t *testing.T) {
	tt := []struct {
		name      string
		in        string
		expectErr bool
	}{
		{
			name:      "valid full cron schedule",
			in:        "* * * * *",
			expectErr: false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			err := parseSchedule(test.in)
			if test.expectErr && err == nil {
				t.Errorf("expected error\ngot no error")
			}
			if !test.expectErr && err != nil {
				t.Errorf("expected no error\ngot error %v", err)
			}
		})
	}
}
