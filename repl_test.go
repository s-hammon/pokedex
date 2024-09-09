package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{
			input: " ",
			want:  []string{},
		},
		{
			input: " hello  ",
			want:  []string{"hello"},
		},
		{
			input: "  hello  world  ",
			want:  []string{"hello", "world"},
		},
		{
			input: " Hell0 World",
			want:  []string{"hello", "world"},
		},
	}

	for _, tt := range cases {
		got := cleanInput(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("cleanInput(%q) = %q; want %q", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("cleanInput(%q) = %q; want %q", tt.input, got, tt.want)
			}
		}
	}
}
