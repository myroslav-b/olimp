package connectors

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRequestOk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`many bytes`))
		return
	}))

	u, _ := url.Parse(server.URL)
	testEdbo := TEdbo{
		client: server.Client(),
		url:    *u,
	}

	result, err := request(&testEdbo)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, []byte(`many bytes`)) {
		t.Error("Bad result")
	}

}

func TestRequestNotOk500(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}))

	u, _ := url.Parse(server.URL)
	testEdbo := TEdbo{
		client: server.Client(),
		url:    *u,
	}

	result, err := request(&testEdbo)

	if err == nil {
		t.Fail()
	}

	if err.Error() != strings.Join([]string{strconv.Itoa(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError)}, " ") {
		t.Fail()
		//t.Error(http.StatusText(http.StatusInternalServerError))
		//t.Error(err.Error())
	}

	if result != nil {
		t.Fail()
	}
}

func TestRequestNotOk404(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		return
	}))

	u, _ := url.Parse(server.URL)
	testEdbo := TEdbo{
		client: server.Client(),
		url:    *u,
	}

	result, err := request(&testEdbo)

	if err == nil {
		t.Fail()
	}

	if err.Error() != strings.Join([]string{strconv.Itoa(http.StatusNotFound), http.StatusText(http.StatusNotFound)}, " ") {
		t.Fail()
		//t.Error(http.StatusText(http.StatusNotFound))
		//t.Error(err.Error())
	}

	if result != nil {
		t.Fail()
	}
}

func TestRequestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		//w.WriteHeader(http.StatusNoContent)
		return
	}))

	u, _ := url.Parse(server.URL)
	testEdbo := TEdbo{
		client: server.Client(),
		url:    *u,
	}

	testEdbo.client.Timeout = 10 * time.Millisecond

	result, err := request(&testEdbo)

	if err == nil {
		t.Fail()
	}

	/*
		if err.Error() != http.StatusText(http.StatusRequestTimeout) {
			//t.Fail()
			t.Error(http.StatusText(http.StatusRequestTimeout))
			t.Error(err)
		}
	*/

	if result != nil {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		have struct {
			bytes []byte
			tag   map[string]string
		}
		want struct {
			batch []map[string]string
			err   error
		}
	}{
		{ //convenient data
			struct {
				bytes []byte
				tag   map[string]string
			}{
				[]byte(`
				[
					{
						"university_name": "Івано-Франківський державний коледж технологій та бізнесу",
						"university_id": "1149",
						"university_parent_id": null,
						"university_short_name": ".",
						"university_name_en": "Ivano-Frankivsk State College of technology and business",
						"is_from_crimea": "",
						"registration_year": "1964",
						"university_edrpou": "01566100",
						"university_type_name": "Заклад вищої освіти",
						"education_type_name": "Коледж",
						"university_financing_type_name": "Державна",
						"university_governance_type_name": "Міністерство освіти і науки України",
						"post_index": "76014",
						"koatuu_id": "2610100000",
						"region_name": "Івано-Франківська область",
						"koatuu_name": "Івано-Франківськ",
						"university_address": "вул. Євгена Коновальця, 140",
						"post_index_u": "76014",
						"koatuu_id_u": "2610100000",
						"region_name_u": "Івано-Франківська область",
						"koatuu_name_u": "",
						"university_address_u": "вул. Євгена Коновальця, 140",
						"university_phone": "(0342) 531400",
						"university_email": "dktbinfo@gmail.com",
						"university_site": "www.dktb.if.ua",
						"university_director_post": "Директор",
						"university_director_fio": "Бабінець Василь Михайлович",
						"close_date": null,
						"primitki": ""
					},
					{
						"university_name": "Івано-Франківський коледж Львівського національного аграрного університету",
						"university_id": "528",
						"university_parent_id": "162",
						"university_short_name": "ІФКЛНАУ",
						"university_name_en": "Ivano-Frankivsk College of Lviv National Agrarian University",
						"is_from_crimea": "",
						"registration_year": "1965",
						"university_edrpou": "35171187",
						"university_type_name": "Заклад вищої освіти",
						"education_type_name": "Коледж",
						"university_financing_type_name": "Державна",
						"university_governance_type_name": "Міністерство освіти і науки України",
						"post_index": "76492",
						"koatuu_id": "2610100000",
						"region_name": "Івано-Франківська область",
						"koatuu_name": "Івано-Франківськ",
						"university_address": "вул. Юності, 11",
						"post_index_u": "76492",
						"koatuu_id_u": "2610100000",
						"region_name_u": "Івано-Франківська область",
						"koatuu_name_u": "",
						"university_address_u": "вул. Юності, 11",
						"university_phone": "0342554897,034554714",
						"university_email": "agrarcol@ukr.net",
						"university_site": "www.ifagrarncol.at.ua",
						"university_director_post": "Директор",
						"university_director_fio": "Костюк Богдан Андрійович",
						"close_date": null,
						"primitki": ""
					}
				]`),
				map[string]string{"id": "university_id", "name": "university_name", "region": "region_name", "locality": "koatuu_name", "address": "university_address", "phone": "university_phone", "email": "university_email", "website": "university_site", "boss": "university_director_fio", "status": "primitki"},
			},
			struct {
				batch []map[string]string
				err   error
			}{
				[]map[string]string{
					{"id": "1149", "name": "Івано-Франківський державний коледж технологій та бізнесу", "region": "Івано-Франківська область", "locality": "Івано-Франківськ", "address": "вул. Євгена Коновальця, 140", "phone": "(0342) 531400", "email": "dktbinfo@gmail.com", "website": "www.dktb.if.ua", "boss": "Бабінець Василь Михайлович", "status": ""},
					{"id": "528", "name": "Івано-Франківський коледж Львівського національного аграрного університету", "region": "Івано-Франківська область", "locality": "Івано-Франківськ", "address": "вул. Юності, 11", "phone": "0342554897,034554714", "email": "agrarcol@ukr.net", "website": "www.ifagrarncol.at.ua", "boss": "Костюк Богдан Андрійович", "status": ""},
				},
				nil,
			},
		},
		{ //inconvenient data: / \/ \" ' \n (` - i can't check)
			struct {
				bytes []byte
				tag   map[string]string
			}{
				[]byte(`
				[
					{
						"university_name": "//Івано-Франківський\n 'державний' \"коледж\" технологій та бізнесу",
						"university_id": "1149",
						"university_parent_id": null,
						"university_short_name": ".",
						"university_name_en": "Ivano-Frankivsk State College of technology and business",
						"is_from_crimea": "",
						"registration_year": "1964",
						"university_edrpou": "01566100",
						"university_type_name": "Заклад вищої освіти",
						"education_type_name": "Коледж",
						"university_financing_type_name": "Державна",
						"university_governance_type_name": "Міністерство освіти і науки України",
						"post_index": "76014",
						"koatuu_id": "2610100000",
						"region_name": "Івано-Франківська область",
						"koatuu_name": "Івано-Франківськ",
						"university_address": "вул. Євгена Коновальця, 140",
						"post_index_u": "76014",
						"koatuu_id_u": "2610100000",
						"region_name_u": "Івано-Франківська область",
						"koatuu_name_u": "",
						"university_address_u": "вул. Євгена Коновальця, 140",
						"university_phone": "(0342) 531400",
						"university_email": "dktbinfo@gmail.com",
						"university_site": "https:\/\/www.dktb.if.ua",
						"university_director_post": "Директор",
						"university_director_fio": "Бабінець Василь Михайлович",
						"close_date": null,
						"primitki": ""
					}
				]`),
				map[string]string{"id": "university_id", "name": "university_name", "region": "region_name", "locality": "koatuu_name", "address": "university_address", "phone": "university_phone", "email": "university_email", "website": "university_site", "boss": "university_director_fio", "status": "primitki"},
			},
			struct {
				batch []map[string]string
				err   error
			}{
				[]map[string]string{
					{"id": "1149", "name": "//Івано-Франківський\n 'державний' \"коледж\" технологій та бізнесу", "region": "Івано-Франківська область", "locality": "Івано-Франківськ", "address": "вул. Євгена Коновальця, 140", "phone": "(0342) 531400", "email": "dktbinfo@gmail.com", "website": "https://www.dktb.if.ua", "boss": "Бабінець Василь Михайлович", "status": ""},
				},
				nil,
			},
		},
		{ //no data
			struct {
				bytes []byte
				tag   map[string]string
			}{
				[]byte(``),
				map[string]string{"id": "university_id", "name": "university_name", "region": "region_name", "locality": "koatuu_name", "address": "university_address", "phone": "university_phone", "email": "university_email", "website": "university_site", "boss": "university_director_fio", "status": "primitki"},
			},
			struct {
				batch []map[string]string
				err   error
			}{
				nil,
				nil,
			},
		},
		{ //nil data
			struct {
				bytes []byte
				tag   map[string]string
			}{
				nil,
				map[string]string{"id": "university_id", "name": "university_name", "region": "region_name", "locality": "koatuu_name", "address": "university_address", "phone": "university_phone", "email": "university_email", "website": "university_site", "boss": "university_director_fio", "status": "primitki"},
			},
			struct {
				batch []map[string]string
				err   error
			}{
				nil,
				nil,
			},
		},
	}
	for _, c := range cases {
		got := make([]map[string]string, len(c.want.batch))
		got, _ = parse(c.have.bytes, c.have.tag)
		if !reflect.DeepEqual(got, c.want.batch) {
			t.Errorf("the values of the results do not match\n%v\n%v", got, c.want.batch)
		}
	}
}
