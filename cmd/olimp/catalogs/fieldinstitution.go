package catalogs

type TFieldInst map[string]string

var FieldInst = TFieldInst{
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

//FieldInstByCode return name of field of institution by code
func FieldInstByCode(code string) (string, bool) {
	f, ok := FieldInst[code]
	return f, ok
}

//IsFieldInst returns true if the value exists, otherwise - false
func IsFieldInst(code string) bool {
	_, ok := FieldInst[code]
	return ok
}
