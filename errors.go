package form_validator

import "fmt"

const (
	ERROR_MISSING_VALUE  = "ERROR_MISSING_VALUE"
	ERROR_INCORRECT_TYPE = "ERROR_INCORRECT_TYPE"
	ERROR_FILE_TYPE      = "ERROR_FILE_TYPE"
)

type FieldError struct {
	Name  string
	Error Error
}

// FormErrors type enables the caller to access all the form errors
// from a map indexed by name.
type FormErrors map[string]map[string]string

func missingValueError(name string) string {
	return fmt.Sprintf("Missing value for %s field", name)
}

func incorrectTypeError(fieldType, name string) string {
	return fmt.Sprintf("Expected a value of type %s for %s field", fieldType, name)
}

func fileError(err error) string {
	return fmt.Sprintf("File error: %e", err)
}

func setErrorMessage(f *Field, fileErr error) {
	switch f.Error.Type {
	case ERROR_MISSING_VALUE:
		f.Error.Message = missingValueError(f.Name)
		break
	case ERROR_INCORRECT_TYPE:
		f.Error.Message = incorrectTypeError(f.Type, f.Name)
		break
	case ERROR_FILE_TYPE:
		f.Error.Message = fileError(fileErr)
	default:
		// pass
	}
}

// GetFormError access a single form error value
//
//		name := GetFormError("name", &c)
//
func GetFormError(name string, c *Config) Error {
	var err Error
	for _, v := range c.Fields {
		if v.Name == name {
			err = v.Error
		}
	}
	return err
}

// GetFormErrors access all form errors as a map (`FormErrors`) indexed off the form names
//
//		var formErrs = form_validator.FormErrors{}
//		form_validator.GetFormErrors(&c, &formErrs)
//
// If the results of `formErrs` are passed to the template as data then
// all form errors can be accessed from the map via index name, for example:
//
//		{{ if .FormErrors.title }}
//                <div class="alert alert-danger" role="alert">
//                    {{ .FormErrors.title.error }}
//                </div>
//      {{ end }}
//
// In this case `FormErrors.title.error` will produce an error message that
// can be safely displayed to the user.
func GetFormErrors(c *Config, fe *FormErrors) {
	for _, v := range c.Fields {
		if v.Error.Type != "" {
			(*fe)[v.Name] = map[string]string{
				v.Name:  v.Name,
				"error": v.Error.Message,
			}
		}
	}
}
