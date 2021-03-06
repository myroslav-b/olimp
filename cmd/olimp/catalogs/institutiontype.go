//Package catalogs contains basic entities and functions for working with them
package catalogs

//TInstType describes type of the institution
type TInstType map[string]string

var instType = TInstType{
	"1": "Заклад вищої освіти",
	"2": "Заклад професійної (професійно-технічної) освіти",
	"3": "Заклад загальної середньої освіти",
	"5": "Орган управління освітою",
	"9": "Заклад фахової передвищої освіти",
}

//InstitutionTypeByCode return type of institution by code
func InstitutionTypeByCode(code string) (string, bool) {
	t, ok := instType[code]
	return t, ok
}

//IsInstitutionType returns true if the value exists, otherwise - false
func IsInstitutionType(code string) bool {
	_, ok := instType[code]
	return ok
}

//MapInstType return the map institution type
func MapInstType() TInstType {
	return instType
}
