package config

import "fmt"

func ErrProfileNotFound(name string) error {
	return fmt.Errorf("profile '%s' not found", name)
}

func ErrGameNotSupported(name string) error {
	return fmt.Errorf("game '%s' is not supported", name)
}
