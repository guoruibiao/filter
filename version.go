package main

import (
	"errors"
)

func init() {
	Register("version", "==", versionEQ)
	Register("version", "!=", versionNE)
}

func versionEQ(input interface{}, operator string, value interface{}) (bool, error) {
	version, ok := value.(string)
	if !ok {
		return false, errors.New("invalid version type")
	}
	inputVersion, ok := input.(string)
	if ok && inputVersion == version {
		return true, nil
	}
	return false, errors.New("version not equal")
}

func versionNE(input interface{}, operator string, value interface{}) (bool, error) {
	version, ok := value.(string)
	if !ok {
		return false, errors.New("invalid version type")
	}

	inputVersion, ok := input.(string)
	if !ok {
		return false, errors.New("invalid inputVersion type")
	}

	if inputVersion != version {
		return true, nil
	}
	return false, errors.New("version equal")
}
