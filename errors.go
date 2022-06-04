package form_validator

import "fmt"

const (
	ERROR_MISSING_VALUE  = "ERROR_MISSING_VALUE"
	ERROR_INCORRECT_TYPE = "ERROR_INCORRECT_TYPE"
)

type FieldError struct {
	Name  string
	Error Error
}

type FormErrors []FieldError

func missingValueError(name string) string {
	return fmt.Sprintf("Missing value for %s field", name)
}

func incorrectTypeError(fieldType, name string) string {
	return fmt.Sprintf("Expected a value of type %s for %s field", fieldType, name)
}

func SetErrorMessage(f *Field) {
	switch f.Error.Type {
	case ERROR_MISSING_VALUE:
		f.Error.Message = missingValueError(f.Name)
		break
	case ERROR_INCORRECT_TYPE:
		f.Error.Message = incorrectTypeError(f.Type, f.Name)
		break
	default:
		// pass
	}
}
