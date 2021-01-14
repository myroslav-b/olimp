package institutions

import "time"

type TRequestType map[string]string

var RequestType = TRequestType{
	"":            "nothing",
	"equal":       "equal",
	"notequal":    "not equal",
	"contains":    "contains",
	"notcontains": "not contains",
	"empty":       "empty",
	"notempty":    "not empty",
}

type TRequest struct {
	V struct {
		InstType string
		RegCode  string
		TInstitutionBundle
	}
	T struct {
		InstType string
		RegCode  string
		TInstitutionBundle
	}
}

type tInstitutionStore map[tCodeBatch]struct {
	shelfLife          time.Duration
	institutionBatches tInstitutionBatch
}

func Query(req TRequest) (tInstitutionBundles, error) {

	return nil, nil
}

/*
func (store tInstitutionStore) add(instType, regCode string, shelfLife time.Duration) error {
	codeBatch := tCodeBatch{instType, regCode}
	return nil
}
*/
func (store tInstitutionStore) del(instType, regCode string) error {

	return nil
}
