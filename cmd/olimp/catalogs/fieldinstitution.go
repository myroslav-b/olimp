//Package catalogs contains basic entities and functions for working with them
package catalogs

//TFieldInst describes map fields of the institution
type TFieldInst map[string]string

var fieldInst = TFieldInst{
	"id":      "унікальний ідентифікатор установи",
	"name":    "назва установи",
	"region":  "регіон (область, Київ, Севастополь, АР Крим) знаходження установи",
	"address": "адреса установи",
	"phone":   "телефон",
	"email":   "адреса електронної пошти",
	"website": "веб-сайт",
	"boss":    "керівник установи",
	"status":  "статус установи: пусто, якщо діюча, інакше - причина припинення діяльності",
}

//MapFieldInst return the map field of institution
func MapFieldInst() TFieldInst {
	return fieldInst
}

//FieldInstByCode return name of field of institution by code
func FieldInstByCode(code string) (string, bool) {
	f, ok := fieldInst[code]
	return f, ok
}

//IsFieldInst returns true if the value exists, otherwise - false
func IsFieldInst(code string) bool {
	_, ok := fieldInst[code]
	return ok
}
