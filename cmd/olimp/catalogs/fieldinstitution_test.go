package catalogs

import (
	"reflect"
	"testing"
)

func TestMapEdboFieldTags(t *testing.T) {
	cases := []struct {
		have string
		want map[string]string
	}{
		{
			"1",
			map[string]string{"id": "university_id", "name": "university_name", "region": "region_name", "locality": "koatuu_name", "address": "university_address", "phone": "university_phone", "email": "university_email", "website": "university_site", "boss": "university_director_fio", "status": "primitki"},
		},
		{
			"3",
			map[string]string{"id": "institution_id", "name": "institution_name", "region": "region_name", "locality": "koatuu_name", "address": "address", "phone": "phone", "email": "email", "website": "website", "boss": "boss", "status": "state_name"},
		},
	}

	for _, c := range cases {
		var got map[string]string
		got = MapEdboFieldTags(c.have)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MapEdboFieldTags(%s) gives \n %v \n want \n %v", c.have, got, c.want)
		}
	}
}
