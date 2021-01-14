package institutions

import "strings"

type TFieldBundle string

type TInstitutionBundle struct {
	id      TFieldBundle
	name    string
	region  string
	address string
	phone   string
	email   string
	website string
	boss    string
	status  string
}

func (field TFieldBundle) checkMatch(st string) bool {
	return st == string(field)
}

func (field TFieldBundle) checkPartMatch(st string) bool {
	return strings.Contains(string(field), st)
}

/*
func (bandle *TInstitutionBundle) checkMatchID(id string) bool {
	return id == bandle.id
}

func (bandle *TInstitutionBundle) checkMatchName(name string) bool {
	return name == bandle.name
}

func (bandle *TInstitutionBundle) checkMatchPartName(str string) bool {
	return strings.Contains(bandle.name, str)
}
*/
