package institutions

import "errors"

var ErrInstType = errors.New("Invalid Institution Type")
var ErrRegCode = errors.New("Invalid Region Code")
var ErrKeyBatch = errors.New("No batch for this key")

var institutionStore tInstitutionStore

func init() {
	institutionStore = make(tInstitutionStore)
}
