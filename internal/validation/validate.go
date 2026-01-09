// package validation

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// )

// var validate = validator.New()

// func Validate(i any) error {
// 	err := validate.Struct(i)
// 	if err != nil {
// 		var errMessages []string

// 		for _, err := range err.(validator.ValidationErrors) {

// 			errMessages = append(errMessages, fmt.Sprintf("%s is required", strings.ToLower(err.Field())))
// 		}

// 		return fmt.Errorf("%s", strings.Join(errMessages, ","))
// 	}
// 	return nil
// }

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
			switch err.Tag() {
			case "required":
				errMessages = append(errMessages, fmt.Sprintf("%s is required", strings.ToLower(err.Field())))
			case "email":
				errMessages = append(errMessages, fmt.Sprintf("%s is not a valid email", strings.ToLower(err.Field())))
			case "min":
				errMessages = append(errMessages, fmt.Sprintf("%s must be at least %s characters", strings.ToLower(err.Field()), err.Param()))
			case "max":
				errMessages = append(errMessages, fmt.Sprintf("%s must be at most %s characters", strings.ToLower(err.Field()), err.Param()))
			default:
				errMessages = append(errMessages, fmt.Sprintf("%s is invalid", strings.ToLower(err.Field())))
			}
		}

		return fmt.Errorf("%s", strings.Join(errMessages, ", "))
	}
	return nil
}
