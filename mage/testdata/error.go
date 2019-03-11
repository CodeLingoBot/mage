// +build mage

package main

import "errors"

// ReturnsNonNilError returns a non-nil error.
func ReturnsNonNilError() error {
	return errors.New("bang!")
}
