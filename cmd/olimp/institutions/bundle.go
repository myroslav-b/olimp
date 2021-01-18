package institutions

import (
	"olimp/cmd/olimp/catalogs"
	"strings"
)

type TInstitutionBundle map[string]string

func (bundle TInstitutionBundle) checkMatchField(field, st string) (bool, error) {
	if _, ok := catalogs.FieldInst[field]; !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := value == st
	return b, nil
}

func (bundle TInstitutionBundle) checkMatchPartField(field, st string) (bool, error) {
	if _, ok := catalogs.FieldInst[field]; !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := strings.Contains(value, st)
	return b, nil
}

func (bundle TInstitutionBundle) checkEmptyField(field string) (bool, error) {
	if _, ok := catalogs.FieldInst[field]; !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := value == ""
	return b, nil
}
