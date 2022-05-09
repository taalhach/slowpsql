package common

import uuid "github.com/nu7hatch/gouuid"

func NewGuid() (string, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return guid.String(), nil
}
