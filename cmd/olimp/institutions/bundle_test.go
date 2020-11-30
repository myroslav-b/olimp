package institutions

import "testing"

func TestCheckContainsName(t *testing.T) {
	bundle := tInstitutionBandle{
		"", "Ново-Нар'ямпільська школа І-ІІІ ступенів №13", "", "", "", "", "", "",
	}

	cases := []struct {
		have string
		want bool
	}{
		{"Ново-Нар'ямпільська школа І-ІІІ ступенів №13", true},
		{"ямпільська", true},
		{"'", true},
		{" ", true},
		{"", true},
		{"Ново-Нар'ямпільська школа І-ІІІ ступенів №13 ", false},
		{"ново-Нар'ямпільська школа І-ІІІ ступенів №13", false},
		{"Ямпільська", false},
		{"  ", false},
	}

	for _, c := range cases {
		got := bundle.checkContainsName(c.have)
		if got != c.want {
			t.Errorf("checkContainsName(%s) == %t, want %t", c.have, got, c.want)
		}

	}
}
