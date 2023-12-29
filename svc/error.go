package svc

import "fmt"

func wrapError(errContext string, err error) error {
	return fmt.Errorf("%s: %w", errContext, err)
}
