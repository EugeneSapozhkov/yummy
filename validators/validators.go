package validators

import (
	"github.com/pkg/errors"
	"regexp"
)

var (
	BadFormatEmail = errors.New("Invalid email")
	BadFormatName = errors.New("Invalid name")
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)


func EmailValid(email string) error {
	if !emailRegexp.MatchString(email) {
		return BadFormatEmail
	}
	return nil
}

func NameValid(name string) error {
	if name == "" {
		return BadFormatName
	}

	return nil
}