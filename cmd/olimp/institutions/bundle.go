//Package institutions contains  logic of request for a list of institutions
package institutions

import (
	"strings"

	"github.com/myroslav-b/olimp/cmd/olimp/catalogs"
)

type tInstitutionBundle map[string]string

func (bundle tInstitutionBundle) checkMatchField(field, st string) (bool, error) {
	if _, ok := catalogs.FieldInstByCode(field); !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := value == st
	return b, nil
}

func (bundle tInstitutionBundle) checkMatchPartField(field, st string) (bool, error) {
	if _, ok := catalogs.FieldInstByCode(field); !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := strings.Contains(value, st)
	return b, nil
}

func (bundle tInstitutionBundle) checkEmptyField(field string) (bool, error) {
	if _, ok := catalogs.FieldInstByCode(field); !ok {
		return false, ErrBadField
	}
	value, ok := bundle[field]
	if !ok {
		return false, ErrNoField
	}
	b := value == ""
	return b, nil
}
