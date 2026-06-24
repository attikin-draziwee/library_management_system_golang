package utils

import (
	"errors"
	"regexp"
	"strings"
)

// Validation Error constant (Iota-like)
const (
	ErrIsEmpty           = "Email пустой"
	ErrInvalidEmail      = "Неверный формат Email"
	ErrHasSpaces         = "в Email должны отсутствовать пробелы"
	ErrLocalTooLong      = "локальная часть > 64 bytes"
	ErrDomainTooLong     = "доменная часть > 255 bytes"
	ErrNonASCIICharacter = "Email должен содержать только ASCII"
)

const emailPattern string = `^[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*@[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)+$`
const nonASCIIPattern string = `^[\x00-\x7F]+$`

// IsValidateEmail validate email address.
// Email specification: https://en.wikipedia.org/wiki/Email_address#RFC_specification
//
// Params:
//   - email: string literal of email
//
// Return:
//
//   - bool - if it is true, err will be always nil
//
//   - err - validation error, regexp error.
//     err use Iota const messages like ErrIsEmpty, ErrHasSpaces and etc.
//
// Examples:
//
//	IsValidateEmail("example@mysite.com") // -> true, nil
//	IsValidateEmail("  ..wrong100%@@ebobo...com") // -> false, err("invalid email format")
func IsValidateEmail(email string) (matched bool, err error) {
	var emailLower string = strings.ToLower(email)
	var parts []string = strings.Split(emailLower, "@")

	switch {
	case strings.TrimSpace(emailLower) == "":
		return false, errors.New(ErrIsEmpty)
	case len(strings.Fields(emailLower)) != 1:
		return false, errors.New(ErrHasSpaces)
	case len(parts) != 2:
		return false, errors.New(ErrInvalidEmail)
	}

	localLen, domainLen := len(parts[0]), len(parts[1])
	switch {
	case localLen > 64:
		return false, errors.New(ErrLocalTooLong)
	case domainLen > 255:
		return false, errors.New(ErrDomainTooLong)
	}

	ASCIIMatched, ASCIIError := regexp.Match(nonASCIIPattern, []byte(emailLower))
	switch {
	case ASCIIError != nil:
		return false, ASCIIError
	case !ASCIIMatched:
		return false, errors.New(ErrNonASCIICharacter)
	}

	if matched, err = regexp.Match(emailPattern, []byte(emailLower)); err != nil {
		return false, err
	}

	if !matched {
		return false, errors.New(ErrInvalidEmail)
	}

	return true, nil
}
