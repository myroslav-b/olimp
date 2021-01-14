package institutions

import (
	"olimp/cmd/olimp/catalogs"
	"time"
)

type tCodeBatch struct {
	instType string
	regCode  string
}

type tInstitutionBundles []TInstitutionBundle

type tInstitutionBatch struct {
	timestamp time.Time
	code      tCodeBatch
	bundles   tInstitutionBundles
}

type BatchLoader interface {
	LoadBatch(instType, regCode string) (tInstitutionBundles, error)
}

func (batch *tInstitutionBatch) setTimestamp() {
	batch.timestamp = time.Now()
}

func (batch *tInstitutionBatch) getTimestamp() time.Time {
	return batch.timestamp
}

func (batch *tInstitutionBatch) setCode(instType, regCode string) error {
	if !catalogs.IsInstitutionType(instType) {
		return ErrInstType
	}
	if !catalogs.IsRegionCode(regCode) {
		return ErrRegCode
	}
	batch.code.instType = instType
	batch.code.regCode = regCode
	return nil
}

func (batch *tInstitutionBatch) getCode() (string, string) {
	return batch.code.instType, batch.code.regCode
}

func (batch *tInstitutionBatch) checkFresh(timeFresh time.Duration) bool {
	b := true
	dt := time.Now().Sub(batch.getTimestamp())
	if dt > timeFresh {
		b = false
	}
	return b
}

func (batch *tInstitutionBatch) checkCode(instType, regCode string) bool {
	it, rc := batch.getCode()
	if (it == instType) && (rc == regCode) {
		return true
	}
	return false
}

func (batch *tInstitutionBatch) init(instType, regCode string, b BatchLoader) error {
	if err := batch.setCode(instType, regCode); err != nil {
		//batch.bundles = make(tInstitutionBundles, 0)
		return err
	}
	batch.setTimestamp()
	bundl, err := b.LoadBatch(instType, regCode)
	if err != nil {
		//batch.bundles = make(tInstitutionBundles, 0)
		return err
	}
	batch.bundles = bundl
	return nil
}
