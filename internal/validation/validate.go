package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(i any) error {
	err := validate.Struct(i)
	if err != nil {
		var errMessages []string

		for _, err := range err.(validator.ValidationErrors) {

			errMessages = append(errMessages, fmt.Sprintf("%s is required", strings.ToLower(err.Field())))
		}

		return fmt.Errorf("%s", strings.Join(errMessages, ","))
	}
	return nil
}
