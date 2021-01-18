package institutions

import (
	"olimp/cmd/olimp/catalogs"
	"reflect"
	"testing"
)

type tFakeEdebo map[tCodeBatch]tInstitutionBundles

func setCode(instType, regCode string) (tCodeBatch, error) {
	var code tCodeBatch
	if !catalogs.IsInstitutionType(instType) {
		return code, ErrInstType
	}
	if !catalogs.IsRegionCode(regCode) {
		return code, ErrRegCode
	}
	code.instType = instType
	code.regCode = regCode
	return code, nil
}

func (fakeEdebo tFakeEdebo) LoadBatch(instType, regCode string) (tInstitutionBundles, error) {
	var instBundles tInstitutionBundles
	code, err := setCode(instType, regCode)
	if err != nil {
		return instBundles, err
	}
	instBundles, ok := fakeEdebo[code]
	if !ok {
		return instBundles, ErrKeyBatch
	}
	return instBundles, nil
}

func TestInitInstitutionBatch(t *testing.T) {
	fakeEdebo := tFakeEdebo{
		{"1", "05"}: {
			//{"", "", "", "", "", "", "", "", ""},
			//{"", "", "", "", "", "", "", "", ""},
			//{"", "", "", "", "", "", "", "", ""},
			{"id": "1609", "name": "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "region": "Вінницька область", "address": "Вінниця, вул. Стрілецька, 7", "phone": "(044) 529 98 31, (043) 227 23 55", "email": "", "website": "", "boss": "Кравчук Галина Олександрівна", "status": ""},
			{"id": "1464", "name": "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "region": "Вінницька область", "address": "Вінниця, вул. Козицького, 15", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			{"id": "1527", "name": "Барська філія Глухівського національного педагогічного університету імені Олександра Довженка", "region": "Вінницька область", "address": "Бар, майдан Михайла Грушевського, 1", "phone": "(043) 412-32-70", "email": "barfiliynpu@ukr.net", "website": "http://new.gnpu.edu.ua", "boss": "Поберецька Вікторія Василівна", "status": ""},
		},
		{"3", "73"}: {
			{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
			{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
		},
		{"2", "01"}: {},
		{"6", "05"}: {},
		{"2", "99"}: {},
	}

	type tHave struct {
		instType string
		regCode  string
	}

	type tWant struct {
		institutionBundles tInstitutionBundles
		err                error
	}

	cases := []struct {
		have tHave
		want tWant
	}{

		{
			tHave{"1", "05"},
			tWant{
				tInstitutionBundles{
					{"id": "1609", "name": "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "region": "Вінницька область", "address": "Вінниця, вул. Стрілецька, 7", "phone": "(044) 529 98 31, (043) 227 23 55", "email": "", "website": "", "boss": "Кравчук Галина Олександрівна", "status": ""},
					{"id": "1464", "name": "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "region": "Вінницька область", "address": "Вінниця, вул. Козицького, 15", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
					{"id": "1527", "name": "Барська філія Глухівського національного педагогічного університету імені Олександра Довженка", "region": "Вінницька область", "address": "Бар, майдан Михайла Грушевського, 1", "phone": "(043) 412-32-70", "email": "barfiliynpu@ukr.net", "website": "http://new.gnpu.edu.ua", "boss": "Поберецька Вікторія Василівна", "status": ""},
				},
				nil,
			},
		},
		{
			tHave{"3", "73"},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			tHave{"2", "01"},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			tHave{"5", "05"},
			tWant{
				//tInstitutionBundles{},
				nil,
				ErrKeyBatch,
			},
		},
		{
			tHave{"6", "05"},
			tWant{
				//tInstitutionBundles{},
				nil,
				ErrInstType,
			},
		},
		{
			tHave{"2", "99"},
			tWant{
				//tInstitutionBundles{},
				nil,
				ErrRegCode,
			},
		},
	}

	for _, c := range cases {
		var gotBatch tInstitutionBatch
		gotErr := gotBatch.init(c.have.instType, c.have.regCode, fakeEdebo)
		if !reflect.DeepEqual(gotBatch.bundles, c.want.institutionBundles) {
			t.Errorf("init(%s, %s, fakeEdebo) gives \n bundles == \n %v \n error == \n %v \n want \n bundles == \n %v \n error == \n %v", c.have.instType, c.have.regCode, gotBatch.bundles, gotErr, c.want.institutionBundles, c.want.err)
		}
		if gotErr != c.want.err {
			t.Errorf("init(%s, %s, fakeEdebo) gives \n bundles == \n %v \n error == \n %v \n want \n bundles == \n %v \n error == \n %v", c.have.instType, c.have.regCode, gotBatch.bundles, gotErr, c.want.institutionBundles, c.want.err)
		}
		/*switch {
		/*case (gotErr != nil) && (c.want.err == nil):
			t.Errorf("init(%s, %s, fakeEdebo) returns %t, want nil", c.have.instType, c.have.regCode, gotErr)
		case (gotErr == nil) && (c.want.err != nil):
			t.Errorf("init(%s, %s, fakeEdebo) returns nil, want %t", c.have.instType, c.have.regCode, c.want.err)
		case (!reflect.DeepEqual(gotBatch.bundles, c.want.institutionBundles)) || (gotErr.Error() != c.want.err.Error()):
			t.Errorf("init(%s, %s, fakeEdebo) gives \n bundles == \n %v \n error == \n %v \n want \n bundles == \n %v \n error == \n %v", c.have.instType, c.have.regCode, gotBatch.bundles, gotErr, c.want.institutionBundles, c.want.err)
		}*/
	}
}
