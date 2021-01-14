package catalogs

import "testing"

type tRez struct {
	ok  bool
	err error
}

func TestCheckByTypeRequest(t *testing.T) {
	cases := []struct {
		have []string
		want tRez
	}{
		{[]string{"abracadabra", "equal", "abracadabra"}, tRez{true, nil}},
		{[]string{"abracadabra", "notequal", "abracadabra"}, tRez{false, nil}},
		{[]string{"abracadabra", "contains", "abracadabra"}, tRez{true, nil}},
		{[]string{"abracadabra", "notcontains", "abracadabra"}, tRez{false, nil}},
		{[]string{"abracadabra", "empty"}, tRez{false, nil}},
		{[]string{"abracadabra", "notempty"}, tRez{true, nil}},
		{[]string{"a \"bra\"  \ncadabra", "equal", "a \"bra\"  \ncadabra"}, tRez{true, nil}},
		{[]string{"a \"bra\"  \ncadabra", "notequal", "a \"bra\"  \ncadabra"}, tRez{false, nil}},
		{[]string{"a \"bra\"  \ncadabra", "contains", " \"bra\"  \ncadabra"}, tRez{true, nil}},
		{[]string{"a \"bra\"  \ncadabra", "notcontains", "a  \"bra\"  \ncadabra"}, tRez{true, nil}},
		{[]string{"a \"bra\"  \ncadabra", "empty"}, tRez{false, nil}},
		{[]string{"a \"bra\"  \ncadabra", "notempty"}, tRez{true, nil}},
		{[]string{"", "equal", ""}, tRez{true, nil}},
		{[]string{"", "notequal", ""}, tRez{false, nil}},
		{[]string{"", "contains", ""}, tRez{true, nil}},
		{[]string{"", "notcontains", ""}, tRez{false, nil}},
		{[]string{"", "empty"}, tRez{true, nil}},
		{[]string{"", "notempty"}, tRez{false, nil}},
		{[]string{"", "equal", " "}, tRez{false, nil}},
		{[]string{" ", "notequal", ""}, tRez{true, nil}},
		{[]string{"", "contains", " "}, tRez{false, nil}},
		{[]string{" ", "notcontains", "  "}, tRez{true, nil}},
		{[]string{" ", "empty"}, tRez{false, nil}},
		{[]string{"  ", "notempty"}, tRez{true, nil}},
		{[]string{"abracadabra", "equal"}, tRez{false, ErrNumberParam}},
		{[]string{"abracadabra", "empty", "abracadabra"}, tRez{false, ErrNumberParam}},
		{[]string{"abracadabra", "abracadabra", "abracadabra"}, tRez{false, ErrTypeRequest}},
		{[]string{"abracadabra", "abracadabra", "abracadabra", "abrakadabra"}, tRez{false, ErrNumberParam}},
		{[]string{"abracadabra"}, tRez{false, ErrNumberParam}},
	}
	var gotOk bool
	var gotErr error
	for _, c := range cases {
		switch len(c.have) {
		case 1:
			gotOk, gotErr = CheckByTypeRequest(c.have[0])
			if (gotOk != c.want.ok) || (gotErr != c.want.err) {
				t.Errorf("checkByTypeRequest(%s) == %t, %s, want %t, %s", c.have[0], gotOk, gotErr, c.want.ok, c.want.err)
			}
		case 2:
			gotOk, gotErr = CheckByTypeRequest(c.have[0], c.have[1])
			if (gotOk != c.want.ok) || (gotErr != c.want.err) {
				t.Errorf("checkByTypeRequest(%s,%s) == %t, %s, want %t, %s", c.have[0], c.have[1], gotOk, gotErr, c.want.ok, c.want.err)
			}
		case 3:
			gotOk, gotErr = CheckByTypeRequest(c.have[0], c.have[1], c.have[2])
			if (gotOk != c.want.ok) || (gotErr != c.want.err) {
				t.Errorf("checkByTypeRequest(%s,%s,%s) == %t, %s, want %t, %s", c.have[0], c.have[1], c.have[2], gotOk, gotErr, c.want.ok, c.want.err)
			}
		case 4:
			gotOk, gotErr = CheckByTypeRequest(c.have[0], c.have[1], c.have[2], c.have[3])
			if (gotOk != c.want.ok) || (gotErr != c.want.err) {
				t.Errorf("checkByTypeRequest(%s,%s,%s,%s) == %t, %s, want %t, %s", c.have[0], c.have[1], c.have[2], c.have[2], gotOk, gotErr, c.want.ok, c.want.err)
			}
		}
	}
}
