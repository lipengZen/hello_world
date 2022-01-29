package error

import (
	// "errors"
	"github.com/pkg/errors"
)

func genErr() error {
	return errors.New("err test")
}

var RecordNotFound = errors.New("RecordNotFound")

func dao() error {
	return RecordNotFound
}

func biz() error {
	err := dao()
	if err != nil {
		return errors.Wrap(err, "wrong sql from biz")
	}

	return nil
}

func biz2222() error {
	err := biz()
	if err != nil {
		return errors.WithMessage(err, "wrong sql from biz2222")
	}

	return nil
}
