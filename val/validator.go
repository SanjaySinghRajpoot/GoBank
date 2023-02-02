package val

import (
	"fmt"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidEmail    = regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9.-]+$`).MatchString
)

func ValidateSting(s string, maxlen int, minlen int) error {
	n := len(s)
	if n > maxlen || n < minlen {
		return fmt.Errorf("must contain %d-%d", minlen, maxlen)
	}

	return nil
}

func ValidateUserName(value string) error {
	if err := ValidateSting(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain letter, digits and underscore")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateSting(value, 5, 15)
}

func ValidateEmail(value string) error {
	if err := ValidateSting(value, 5, 15); err != nil {
		return err
	}

	if !isValidEmail(value) {
		return fmt.Errorf("invalid email address")
	}

	return nil
}
