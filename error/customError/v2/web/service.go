package web

import (
	"goTestProject/error/customError/v2/errors"
)

func getBlogService(id int) (Model, error) {
	result, err := getBlog(id)
	if err != nil {
		switch err {
		case NoResult:
			err = errors.MySQLNoQueryError.Wrap(err)
		case NoField:
			err = errors.MySQLNoFieldError.Wrap(err)
		}

		return nil, err
	}
	return result, nil
}
