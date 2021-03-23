//Package institutions contains  logic of request for a list of institutions
package institutions

import (
	"errors"
	"olimp/cmd/olimp/connectors"
	"time"
)

//ErrInstType - error "Invalid Institution Type"
var ErrInstType = errors.New("Invalid Institution Type")

//ErrRegCode - error "Invalid Region Code"
var ErrRegCode = errors.New("Invalid Region Code")

//ErrKeyBatch - error "No batch for this key"
var ErrKeyBatch = errors.New("No batch for this key")

//ErrBadField - error "Bad field of institution"
var ErrBadField = errors.New("Bad field of institution")

//ErrNoField - error "No field of institution"
var ErrNoField = errors.New("No field of institution")

var institutionStore tInstitutionStore

//ext init
//var batchLoader iBatchLoader = nil
var nameLoader = "edbo"
var tempShelfLife = time.Duration(24 * 60 * 60 * 1000000000)

func init() {
	institutionStore.institutionBatches = make(tInstitutionBatches)
	institutionStore.setShelfLife(tempShelfLife)
	switch nameLoader {
	case "edbo":
		//institutionStore.setBatchLoader(connectors.NewEdboLoader())
		//var conn connectors.TEdboLoader
		//institutionStore.batchLoader = connectors.EdboLoader
		institutionStore.batchLoader = connectors.NewEdboLoader()
	}

	institutionStore.initTurnstile()
}
