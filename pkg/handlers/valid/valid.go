package valid

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func Value(value string) error {
	if !govalidator.IsNumeric(value) {
		return errors.New(value + " is not numeric")
	}
	return nil
}
