package web

import (
	"fmt"
	"goTestProject/error/customError/errors"
)

func getResultService(id int) (Model, error) {
	result, err := getResult(id)
	if err != nil {
		msg := fmt.Sprintf("error getting the  result with id %d", id)
		switch err {
		case NoResult:
			err = errors.NotFound.Wrap(err, msg)
		default:
			err = errors.Wrapf(err, msg)
		}
		return nil, err
	}
	return result, nil
}
