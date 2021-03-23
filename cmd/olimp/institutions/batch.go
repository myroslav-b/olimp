//Package institutions contains  logic of request for a list of institutions
package institutions

import (
	"olimp/cmd/olimp/catalogs"
	"time"
)

type tCodeBatch struct {
	instType string
	regCode  string
}

//tInstitutionBundles describes the bundle of institutions
type tInstitutionBundles []map[string]string

//type tInstitutionBundles []tInstitutionBundle

type tInstitutionBatch struct {
	timestamp time.Time
	code      tCodeBatch
	bundles   tInstitutionBundles
}

type iBatchLoader interface {
	LoadBatch(instType, regCode string) ([]map[string]string, error)
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

func (batch *tInstitutionBatch) init(instType, regCode string, b iBatchLoader) error {
	if err := batch.setCode(instType, regCode); err != nil {
		return err
	}
	batch.setTimestamp()
	bundles, err := b.LoadBatch(instType, regCode)
	if err != nil {
		return err
	}
	batch.bundles = tInstitutionBundles(bundles)
	return nil
}
