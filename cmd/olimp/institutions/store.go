//Package institutions contains  logic of request for a list of institutions
package institutions

import (
	"time"

	"github.com/myroslav-b/olimp/cmd/olimp/catalogs"
	"github.com/myroslav-b/olimp/cmd/olimp/connectors"
)

/*
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
*/

//TRequest describes format of request to obtain a list of institutions
type TRequest struct {
	instType  string
	regCode   string
	valFields tInstitutionBundle
	reqFields tInstitutionBundle
}

type tInstitutionBatches map[tCodeBatch]tInstitutionBatch
type tMapChan map[tCodeBatch]chan bool

type tInstitutionStore struct {
	shelfLife          time.Duration
	institutionBatches tInstitutionBatches
	batchLoader        iBatchLoader
	turnstile          tMapChan
}

/*func (insts tInstitutionStore) QueryInstitutions(it, rc string, vf, rf map[string]string) ([]map[string]string, error) {
	instBundles, err := Query(struct {
		instType, regCode    string
		valFields, reqFields tInstitutionBundle
	}{it, rc, vf, rf})
	if err != nil {
		return nil, err
	}
	return instBundles, nil
}*/

func CreateRequest(it, rc string, vf, rf map[string]string) TRequest {
	cr := TRequest{it, rc, vf, rf}
	return cr
}

func InitStore(nameLoader string, tempShelfLife time.Duration) {
	institutionStore.institutionBatches = make(tInstitutionBatches)
	institutionStore.setShelfLife(tempShelfLife)
	switch nameLoader {
	case "edbo":

		institutionStore.batchLoader = connectors.NewEdboLoader()
	}

	institutionStore.initTurnstile()
}

func (store *tInstitutionStore) initTurnstile() {
	var codeBatch tCodeBatch
	store.turnstile = make(tMapChan, (len(catalogs.MapInstType())+1)*(len(catalogs.MapRegCode())+1))
	for instType := range catalogs.MapInstType() {
		for regCode := range catalogs.MapRegCode() {
			codeBatch.instType = instType
			codeBatch.regCode = regCode
			c := make(chan bool, 1)
			store.turnstile[codeBatch] = c
		}
	}
}

func (store *tInstitutionStore) goToTurnstile(codeBatch tCodeBatch) {
	store.turnstile[codeBatch] <- true
}

func (store *tInstitutionStore) getOutTurnstile(codeBatch tCodeBatch) {
	<-store.turnstile[codeBatch]
}

func (store *tInstitutionStore) setBatchLoader(bl iBatchLoader) {
	store.batchLoader = bl
}

func (store *tInstitutionStore) setShelfLife(shelfLife time.Duration) {
	store.shelfLife = shelfLife
}

//Query executes a request for a list of institutions
func Query(req TRequest) (tInstitutionBundles, error) {

	if !catalogs.IsInstitutionType(req.instType) {
		return nil, ErrInstType
	}
	if !catalogs.IsRegionCode(req.regCode) {
		return nil, ErrRegCode
	}
	f := true
	for key := range req.valFields {
		f = f && catalogs.IsFieldInst(key)
	}
	if !f {
		return nil, ErrBadField
	}
	f = true
	for key := range req.reqFields {
		f = f && catalogs.IsFieldInst(key)
	}
	if !f {
		return nil, ErrBadField
	}

	var codeBatch tCodeBatch
	codeBatch.instType = req.instType
	codeBatch.regCode = req.regCode

	institutionStore.goToTurnstile(codeBatch)
	defer institutionStore.getOutTurnstile(codeBatch)

	batch, ok := institutionStore.institutionBatches[codeBatch]
	switch {
	case !ok:
		{
			err := batch.init(req.instType, req.regCode, institutionStore.batchLoader)
			if err != nil {
				return nil, err
			}
			institutionStore.institutionBatches[codeBatch] = batch
		}
	case ok && !batch.checkFresh(institutionStore.shelfLife):
		{
			institutionStore.delBatch(codeBatch)
			err := batch.init(req.instType, req.regCode, institutionStore.batchLoader)
			if err != nil {
				return nil, err
			}
			institutionStore.institutionBatches[codeBatch] = batch
		}
	}

	rezBundles := make(tInstitutionBundles, 0)
	for _, bundle := range batch.bundles {
		flag := true
		for field := range bundle {
			bool, err := catalogs.CheckByTypeRequest(bundle[field], req.reqFields[field], req.valFields[field])
			if err != nil {
				return nil, err
			}
			flag = flag && bool
		}
		if flag {
			rezBundles = append(rezBundles, bundle)
		}
	}

	return rezBundles, nil
}

/*
func (store tInstitutionStore) add(instType, regCode string, shelfLife time.Duration) error {
	codeBatch := tCodeBatch{instType, regCode}
	return nil
}
*/
func (store tInstitutionStore) delBatch(codeBatch tCodeBatch) {
	delete(store.institutionBatches, codeBatch)
	return
}
