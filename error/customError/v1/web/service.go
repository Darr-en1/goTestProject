package web

import (
	"fmt"
	errors2 "goTestProject/error/customError/v1/errors"
)

func getResultService(id int) (Model, error) {
	result, err := getResult(id)
	if err != nil {
		msg := fmt.Sprintf("error getting the  result with id %d", id)
		switch err {
		case NoResult:
			err = errors2.NotFound.Wrap(err, msg)
		default:
			err = errors2.Wrapf(err, msg)
		}
		return nil, err
	}
	return result, nil
}
