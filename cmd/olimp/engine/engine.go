package engine

import "olimp/cmd/olimp/institutions"

type tStore struct{}

func (s tStore) QueryInstitutions(
	instType, regCode string, valFields, reqFields map[string]string) ([]map[string]string, error) {
	insts, err := institutions.Query(institutions.CreateRequest(instType, regCode, valFields, reqFields))
	if err != nil {
		return nil, err
	}
	return insts, nil
}

func CreateStore() tStore {
	return struct{}{}
}
