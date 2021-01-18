package institutions

import "testing"

func TestCheckContainsName(t *testing.T) {
	bundle := TInstitutionBundle{
		"id": "", "name": "Ново-Нар'ямпільська школа І-ІІІ ступенів №13", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": "",
	}

	type tHave struct {
		field string
		st    string
	}

	type tWant struct {
		ok  bool
		err error
	}

	type tGot struct {
		ok  bool
		err error
	}

	var got tGot

	cases := []struct {
		have tHave
		want tWant
	}{
		{tHave{"name", "Ново-Нар'ямпільська школа І-ІІІ ступенів №13"}, tWant{true, nil}},
		{tHave{"name", "ямпільська"}, tWant{true, nil}},
		{tHave{"name", "'"}, tWant{true, nil}},
		{tHave{"name", " "}, tWant{true, nil}},
		{tHave{"name", ""}, tWant{true, nil}},
		{tHave{"name", "Ново-Нар'ямпільська школа І-ІІІ ступенів №13 "}, tWant{false, nil}},
		{tHave{"name", "ново-Нар'ямпільська школа І-ІІІ ступенів №13"}, tWant{false, nil}},
		{tHave{"name", "Ямпільська"}, tWant{false, nil}},
		{tHave{"name", "  "}, tWant{false, nil}},
		{tHave{"inst", "Ново-Нар'ямпільська школа І-ІІІ ступенів №13"}, tWant{false, ErrBadField}},
	}

	for _, c := range cases {
		got.ok, got.err = bundle.checkMatchPartField(c.have.field, c.have.st)
		if (got.ok != c.want.ok) || (got.err != c.want.err) {
			t.Errorf("checkContainsName(%s) == %t, want %t", c.have, got, c.want)
		}

	}
}
