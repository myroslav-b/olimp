package institutions

import (
	"reflect"
	"testing"
	"time"
)

type tWant struct {
	bundles tInstitutionBundles
	err     error
}

type tFakeEdebo2 map[tCodeBatch]([]map[string]string)

//type tFakeEdebo2 map[tCodeBatch]tInstitutionBundles

func (fakeEdebo tFakeEdebo2) LoadBatch(instType, regCode string) ([]map[string]string, error) {
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

func TestQuery(t *testing.T) {
	fakeEdebo := tFakeEdebo2{
		{"1", "05"}: {
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

	//institutionStore.setShelfLife(time.Duration(24 * 60 * 60 * 1000000000))
	institutionStore.setBatchLoader(fakeEdebo)

	InitStore(":8080", time.Duration(24*60*60*1000000000))

	cases := []struct {
		have TRequest
		want tWant
	}{
		{
			TRequest{
				"13", "85",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrInstType,
			},
		},
		{
			TRequest{
				"1", "99",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrRegCode,
			},
		},
		{
			TRequest{
				"1", "85",
				tInstitutionBundle{"id": "", "name": "", "abrakadabra": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrBadField,
			},
		},
		{
			TRequest{
				"1", "85",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "abrakadabra": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrBadField,
			},
		},
		{
			TRequest{
				"1", "85",
				tInstitutionBundle{"abrakadabra": "", "id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrBadField,
			},
		},
		{
			TRequest{
				"1", "85",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "abrakadabra": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrBadField,
			},
		},
		{
			TRequest{
				"1", "85",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				nil,
				ErrKeyBatch,
			},
		},
		//then - interdependent tests, consistency is important
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{},
				tInstitutionBundle{},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				nil,
				nil,
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "123456", "name": "abrakadabra", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				nil,
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				nil,
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				nil,
				tInstitutionBundle{"id": "", "name": "notempty", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
					{"id": "143955", "name": "Годилівський навчально-виховний комплекс Великокучурівської сільської ради Сторожинецького району Чернівецької області", "region": "Чернівецька область", "address": "с. Годилів, Сторожинецький район, Чернівецька область, вул. Центральна, буд.64", "phone": "(0373)569245", "email": "goduliv@ukr.net", "website": "http://goduliv.at.ua", "boss": "Директор Настасійчук Людмила Дмитрівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "135693", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "equal", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "135693", "name": "", "region": ""},
				tInstitutionBundle{"id": "equal", "name": "", "region": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "135693"},
				tInstitutionBundle{"id": "equal", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "135693", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "equal"},
			},
			tWant{
				tInstitutionBundles{
					{"id": "135693", "name": "комунальна обласна спеціалізована школа-інтернат ІІ-ІІІ ступенів з поглибленим вивченням окремих предметів \"Багатопрофільний ліцей для обдарованих дітей\"", "region": "Чернівецька область", "address": "Чернівці, Чернівецька область, вул. Винниченка, буд.119", "phone": "(0372)522142", "email": "oblinter119@ukr.net", "website": "", "boss": "Семанюк Марина Костянтинівна", "status": ""},
				},
				nil,
			},
		},
		{
			TRequest{
				"3", "73",
				tInstitutionBundle{"id": "111111", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "equal", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			TRequest{
				"1", "05",
				tInstitutionBundle{"id": "111111", "name": "університет", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "notequal", "name": "contains", "region": "notempty", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
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
			TRequest{
				"1", "05",
				tInstitutionBundle{"id": "111111", "name": "університет", "region": "xxxxxx", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": "xxxxxx"},
				tInstitutionBundle{"id": "notequal", "name": "contains", "region": "notempty", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": "empty"},
			},
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
			TRequest{
				"2", "01",
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			TRequest{
				"2", "01",
				tInstitutionBundle{"id": "", "name": "університет", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "", "name": "notcontains", "region": "contains", "address": "", "phone": "empty", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			TRequest{
				"2", "01",
				tInstitutionBundle{"id": "123456", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
				tInstitutionBundle{"id": "equal", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			TRequest{
				"2", "01",
				tInstitutionBundle{},
				tInstitutionBundle{},
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
		{
			TRequest{
				"2", "01",
				nil,
				nil,
			},
			tWant{
				tInstitutionBundles{},
				nil,
			},
		},
	}

	for _, c := range cases {
		type tGot struct {
			bundles tInstitutionBundles
			err     error
		}
		var got tGot
		got.bundles, got.err = Query(c.have)
		if !reflect.DeepEqual(c.want.bundles, got.bundles) || (c.want.err != got.err) {
			t.Errorf("Query(%v) \n == \n Bundles = %v, Error = %s \n want \n Bundles = %v, Error = %s", c.have, got.bundles, got.err, c.want.bundles, c.want.err)
		}
	}

	institutionStore.setShelfLife(time.Duration(0))
	tempCode, _ := setCode("1", "05")
	delete(fakeEdebo, tempCode)
	tempHaveBundles := tInstitutionBundles{
		{"id": "1609", "name": "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "region": "Вінницька область", "address": "Вінниця, вул. Стрілецька, 7", "phone": "(044) 529 98 31, (043) 227 23 55", "email": "", "website": "", "boss": "Кравчук Галина Олександрівна", "status": ""},
		{"id": "1464", "name": "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "region": "Вінницька область", "address": "Вінниця, вул. Козицького, 15", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
	}
	tempRequest := TRequest{
		"1", "05",
		tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
		tInstitutionBundle{"id": "", "name": "", "region": "", "address": "", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
	}
	tempWantBundles := tInstitutionBundles{
		{"id": "1609", "name": "Відокремлений підрозділ 'Вінницький факультет Київського національного університету культури і мистецтв'", "region": "Вінницька область", "address": "Вінниця, вул. Стрілецька, 7", "phone": "(044) 529 98 31, (043) 227 23 55", "email": "", "website": "", "boss": "Кравчук Галина Олександрівна", "status": ""},
		{"id": "1464", "name": "Вінницька філія Вищого навчального закладу \"Київський університет ринкових відносин\"", "region": "Вінницька область", "address": "Вінниця, вул. Козицького, 15", "phone": "", "email": "", "website": "", "boss": "", "status": ""},
	}
	fakeEdebo[tempCode] = tempHaveBundles
	tempGotBundles, _ := Query(tempRequest)
	if !reflect.DeepEqual(tempWantBundles, tempGotBundles) {
		t.Errorf("Check fresh batch: Query(%v) \n == \n Bundles = %v \n want \n Bundles = %v", tempRequest, tempGotBundles, tempWantBundles)
	}

}
