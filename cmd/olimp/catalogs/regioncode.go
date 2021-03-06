//Package catalogs contains basic entities and functions for working with them
package catalogs

//TRegCode describes region codes
type TRegCode map[string]string

var regCode = TRegCode{
	"01": "Автономна Республіка Крим",
	"05": "Вінницька область",
	"07": "Волинська область",
	"12": "Дніпропетровська область",
	"14": "Донецька область",
	"18": "Житомирська область",
	"21": "Закарпатська область",
	"23": "Запорізька область",
	"26": "Івано-Франківська область",
	"32": "Київська область",
	"35": "Кіровоградська область",
	"44": "Луганська область",
	"46": "Львівська область",
	"48": "Миколаївська область",
	"51": "Одеська область",
	"53": "Полтавська область",
	"56": "Рівненська область",
	"59": "Сумська область",
	"61": "Тернопільська область",
	"63": "Харківська область",
	"65": "Херсонська область",
	"68": "Хмельницька область",
	"71": "Черкаська область",
	"73": "Чернівецька область",
	"74": "Чернігівська область",
	"80": "м. Київ",
	"85": "м. Севастополь",
}

//RegionNameByCode return name of region by code
func RegionNameByCode(code string) (string, bool) {
	n, ok := regCode[code]
	return n, ok
}

//IsRegionCode returns true if the value exists, otherwise - false
func IsRegionCode(code string) bool {
	_, ok := regCode[code]
	return ok
}

//MapRegCode return the map region code
func MapRegCode() TRegCode {
	return regCode
}
