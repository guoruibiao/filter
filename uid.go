package main

import (
	"errors"
)

func init() {
	Register("uid", "in", uidIn)
	Register("uid", "notin", uidNotIn)
}

func uidIn(input interface{}, operator string, value interface{}) (bool, error) {
	if operator != "in" {
		return false, errors.New("operator should be `in`")
	}

	arr, ok := value.([]interface{})
	if !ok {
		return false, errors.New("invalid uid type")
	}

	inputUID, ok := input.(string)
	if !ok {
		return false, errors.New("invalid inputUid type")
	}

	for idx := 0; idx < len(arr); idx++ {
		if item, _ := arr[idx].(string); item == inputUID {
			return true, nil
		}
	}

	return false, errors.New("inputUID not exists")
}

func uidNotIn(input interface{}, operator string, value interface{}) (bool, error) {
	if operator != "notin" {
		return false, errors.New("operator should be `notin`")
	}

	arr, ok := value.([]string)
	if !ok {
		return false, errors.New("invalid uid type")
	}

	inputUID, ok := input.(string)
	if !ok {
		return false, errors.New("invalid inputUid type")
	}

	for idx := 0; idx < len(arr); idx++ {
		if arr[idx] == inputUID {
			return false, errors.New("inputUID exists")
		}
	}

	return true, nil
}
