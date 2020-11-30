package institutions

import "strings"

type tInstitutionBandle struct {
	id      string
	name    string
	region  string
	address string
	phone   string
	email   string
	website string
	boss    string
}

func (bandle tInstitutionBandle) checkMatchID(id string) bool {
	return id == bandle.id
}

func (bandle tInstitutionBandle) checkMatchName(name string) bool {
	return name == bandle.name
}

func (bandle tInstitutionBandle) checkContainsName(str string) bool {
	return strings.Contains(bandle.name, str)
}
