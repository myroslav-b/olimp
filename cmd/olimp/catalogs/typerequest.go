//Package catalogs contains basic entities and functions for working with them
package catalogs

import (
	"errors"
	"strings"
)

//TTypeRequest describes request type
type TTypeRequest map[string]string

var typeRequest = TTypeRequest{
	"":            "nothing",
	"equal":       "equal",
	"notequal":    "not equal",
	"contains":    "contains",
	"notcontains": "not contains",
	"empty":       "empty",
	"notempty":    "not empty",
}

//ErrTypeRequest - error "Request type does not exist"
var ErrTypeRequest = errors.New("Request type does not exist")

//ErrNumberParam - error "Wrong number of parameters"
var ErrNumberParam = errors.New("Wrong number of parameters")

//MapTypeRequest return the map type request
func MapTypeRequest() TTypeRequest {
	return typeRequest
}

//TypeRequestByCode return type of institution by code
func TypeRequestByCode(code string) (string, bool) {
	t, ok := typeRequest[code]
	return t, ok
}

//IsTypeRequest returns true if the value exists, otherwise - false
func IsTypeRequest(code string) bool {
	_, ok := typeRequest[code]
	return ok
}

//CheckByTypeRequest applies a condition typeRequest to strings (string) as a binary (unary) operator
func CheckByTypeRequest(st ...string) (bool, error) {
	if (len(st) < 2) || (len(st) > 3) {
		return false, ErrNumberParam
	}
	if !IsTypeRequest(st[1]) {
		return false, ErrTypeRequest
	}
	switch st[1] {
	case "":
		if len(st) != 2 && (len(st) != 3) {
			return false, ErrNumberParam
		}
		return true, nil
	case "equal":
		if len(st) != 3 {
			return false, ErrNumberParam
		}
		if st[0] == st[2] {
			return true, nil
		}
	case "notequal":
		if len(st) != 3 {
			return false, ErrNumberParam
		}
		if st[0] != st[2] {
			return true, nil
		}
	case "contains":
		if len(st) != 3 {
			return false, ErrNumberParam
		}
		if strings.Contains(st[0], st[2]) {
			return true, nil
		}
	case "notcontains":
		if len(st) != 3 {
			return false, ErrNumberParam
		}
		if !strings.Contains(st[0], st[2]) {
			return true, nil
		}
	case "empty":
		if len(st) != 2 && (len(st) != 3) {
			return false, ErrNumberParam
		}
		if st[0] == "" {
			return true, nil
		}
	case "notempty":
		if len(st) != 2 && (len(st) != 3) {
			return false, ErrNumberParam
		}
		if st[0] != "" {
			return true, nil
		}
	}
	return false, nil
}
