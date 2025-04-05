package greet

import "errors"

func (s *service) Meet(name, gender string) (msg string, err error) {
	if len(name) < 0 {
		return "", errors.New("empty name")
	}

	if len(gender) < 0 {
		return "", errors.New("empty gender")
	}

	if gender == "male" {
		msg = "Salom akai "
	} else if gender == "female" {
		msg = "Salom apai "
	}

	msg += name

	return msg, nil
}
