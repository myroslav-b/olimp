//Package catalogs contains basic entities and functions for working with them
package catalogs

//tFieldInst describes map fields of the institution
type tFieldInst map[string]string

var fieldInst = tFieldInst{
	"id":       "унікальний ідентифікатор установи",
	"name":     "назва установи",
	"region":   "регіон (область, Київ, Севастополь, АР Крим) знаходження установи",
	"locality": "населений пункт в якому знаходиться установа",
	"address":  "адреса установи",
	"phone":    "телефон",
	"email":    "адреса електронної пошти",
	"website":  "веб-сайт",
	"boss":     "керівник установи",
	"status":   "статус установи: пусто, якщо діюча, інакше - причина припинення діяльності",
}

//TEdboFieldTags describes json tags for fields of the institution
type TEdboFieldTags map[string]map[string]string

var edboFieldTags = TEdboFieldTags{
	"id":       {"1": "university_id", "2": "university_id", "3": "institution_id", "5": "university_id", "9": "university_id"},
	"name":     {"1": "university_name", "2": "university_name", "3": "institution_name", "5": "university_name", "9": "university_name"},
	"region":   {"1": "region_name", "2": "region_name", "3": "region_name", "5": "region_name", "9": "region_name"},
	"locality": {"1": "koatuu_name", "2": "koatuu_name", "3": "koatuu_name", "5": "koatuu_name", "9": "koatuu_name"},
	"address":  {"1": "university_address", "2": "university_address", "3": "address", "5": "university_address", "9": "university_address"},
	"phone":    {"1": "university_phone", "2": "university_phone", "3": "phone", "5": "university_phone", "9": "university_phone"},
	"email":    {"1": "university_email", "2": "university_email", "3": "email", "5": "university_email", "9": "university_email"},
	"website":  {"1": "university_site", "2": "university_site", "3": "website", "5": "university_site", "9": "university_site"},
	"boss":     {"1": "university_director_fio", "2": "university_director_fio", "3": "boss", "5": "university_director_fio", "9": "university_director_fio"},
	"status":   {"1": "primitki", "2": "primitki", "3": "state_name", "5": "primitki", "9": "primitki"},
}

//MapFieldInst return the map field of institution
func MapFieldInst() tFieldInst {
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

//MapEdboFieldTags returns a json tag map for the fields of the institution of the specified type
func MapEdboFieldTags(instType string) map[string]string {
	tags := make(map[string]string, len(fieldInst))
	for field := range edboFieldTags {
		tags[field] = edboFieldTags[field][instType]
	}
	return tags
}
