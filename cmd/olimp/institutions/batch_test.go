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
			{"1609", "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "Вінницька область", "Вінниця, вул. Стрілецька, 7", "(044) 529 98 31, (043) 227 23 55", "", "", "Кравчук Галина Олександрівна", ""},
			{"1464", "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "Вінницька область", "Вінниця, вул. Козицького, 15", "", "", "", "", ""},
			{"1527", "Барська філія Глухівського національного педагогічного університету імені Олександра Довженка", "Вінницька область", "Бар, майдан Михайла Грушевського, 1", "(043) 412-32-70", "barfiliynpu@ukr.net", "http://new.gnpu.edu.ua", "Поберецька Вікторія Василівна", ""},
		},
		{"3", "73"}: {
			{"135693", "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "Чернівецька область", "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "(0372)522142", "oblinter119@ukr.net", "", "Семанюк Марина Костянтинівна", ""},
			{"143955", "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "Чернівецька область", "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "(0373)569245", "goduliv@ukr.net", "http://goduliv.at.ua", "Директор Настасійчук Людмила Дмитрівна", ""},
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
					{"1609", "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "Вінницька область", "Вінниця, вул. Стрілецька, 7", "(044) 529 98 31, (043) 227 23 55", "", "", "Кравчук Галина Олександрівна", ""},
					{"1464", "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "Вінницька область", "Вінниця, вул. Козицького, 15", "", "", "", "", ""},
					{"1527", "Барська філія Глухівського національного педагогічного університету імені Олександра Довженка", "Вінницька область", "Бар, майдан Михайла Грушевського, 1", "(043) 412-32-70", "barfiliynpu@ukr.net", "http://new.gnpu.edu.ua", "Поберецька Вікторія Василівна", ""},
				},
				nil,
			},
		},
		{
			tHave{"3", "73"},
			tWant{
				tInstitutionBundles{
					{"135693", "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "Чернівецька область", "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "(0372)522142", "oblinter119@ukr.net", "", "Семанюк Марина Костянтинівна", ""},
					{"143955", "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "Чернівецька область", "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "(0373)569245", "goduliv@ukr.net", "http://goduliv.at.ua", "Директор Настасійчук Людмила Дмитрівна", ""},
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
