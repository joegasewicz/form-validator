package form_validator

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Config the `Fields` struct field is where the form values are declared
//
//			c := form_validator.Config{
//	   		MaxMemory: 0,
//	   		Fields: []form_validator.Field{
//	       		{
//	           		Name:     "name",
//	           		Validate: false,
//	           		Default:  "John",
//	           		Type:     "string",
//	       		},
//	       		{
//	           		Name:     "email",
//	           		Validate: false,
//	           		Default:  "",
//	           		Type:     "string",
//	       		},
//	   		},
//			}
//
//		The types are as followed
//
// - Name field is the form's 'name' value
// - Validate sets whether the field requires validation
// - Default set a default value is the form field empty
// - Type sets the type conversion e.g. int8, uint, float16 ...
type Config struct {
	MaxMemory int64
	Fields    []Field
}

// Field represents a form field
type Field struct {
	Name     string
	Validate bool
	Default  string
	Type     string
	Value    interface{}
	Initial  string
	Error    Error
	Matches  string
}

// Error object holds the error type & a message to display to the user
type Error struct {
	Message string
	Type    string
}

// ValidateForm validates a form
//
//	if ok := form_validator.ValidateForm(r, &c); ok {
//		// form is valid
//	} else {
//		// form is invalid
//	}
func ValidateForm(r *http.Request, c *Config) bool {
	for _, f := range c.Fields {
		if f.Type == "file" {
			panic("You must use ValidateMultiPartForm function to parse MultiPartForm data")
		}
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	validate(r, c)
	return isFormValid(c)
}

// ValidateMultiPartForm validates a multipart form
//
//	if ok := form_validator.ValidateMultiPartForm(r, &c); ok {
//		// form is valid
//	} else {
//		// form is invalid
//	}
func ValidateMultiPartForm(r *http.Request, c *Config) bool {
	r.ParseMultipartForm(c.MaxMemory)
	validate(r, c)
	return isFormValid(c)
}

func isFormValid(c *Config) bool {
	for _, f := range c.Fields {
		if f.Error.Type != "" {
			return false
		}
	}
	return true
}

func setValueToInitialOrDefault(f *Field) string {
	// Returns value only if first, initial has a value
	// then if default exist otherwise returns ""
	if f.Initial != "" {
		return f.Initial
	}
	if f.Default != "" {
		return f.Default
	}
	return ""
}

func convertToType(f *Field) {
	switch f.Type {
	case "string":
		f.Value = setValueToInitialOrDefault(f)
		break
	case "bool":
		initialOrDefault := setValueToInitialOrDefault(f)
		b, err := strconv.ParseBool(initialOrDefault)
		if err != nil {
			log.Printf("Error converting value of %s to type bool\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = b
	case "float32":
		initialOrDefault := setValueToInitialOrDefault(f)
		float, err := strconv.ParseFloat(initialOrDefault, 32)
		if err != nil {
			log.Printf("Error converting value of %s to type float32\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		} else {
			f.Value = float32(float)
		}
		break
	case "float64":
		initialOrDefault := setValueToInitialOrDefault(f)
		float, err := strconv.ParseFloat(initialOrDefault, 64)
		if err != nil {
			log.Printf("Error converting value of %s to type float64\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		} else {
			f.Value = float // float64
		}
		break
	case "int":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.Atoi(initialOrDefault)
		if err != nil {
			log.Printf("Error converting value of %s to type int\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = u
	case "uint":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseUint(initialOrDefault, 10, 64)
		if err != nil {
			log.Printf("Error converting value of %s to type uint\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = uint(u)
	case "uint8":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseUint(initialOrDefault, 10, 8)
		if err != nil {
			log.Printf("Error converting value of %s to type uint8\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = uint8(u)
	case "uint16":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseUint(initialOrDefault, 10, 16)
		if err != nil {
			log.Printf("Error converting value of %s to type uint16\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = uint16(u)
	case "uint32":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseUint(initialOrDefault, 10, 32)
		if err != nil {
			log.Printf("Error converting value of %s to type uint32\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = uint32(u)
	case "uint64":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseUint(initialOrDefault, 10, 64)
		if err != nil {
			log.Printf("Error converting value of %s to type uint64\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = u // uint64
	case "int8":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseInt(initialOrDefault, 10, 8)
		if err != nil {
			log.Printf("Error converting value of %s to type int8\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = int8(u)
	case "int16":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseInt(initialOrDefault, 10, 16)
		if err != nil {
			log.Printf("Error converting value of %s to type int16\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = int16(u)
	case "int32":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseInt(initialOrDefault, 10, 32)
		if err != nil {
			log.Printf("Error converting value of %s to type int32\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = int32(u)
	case "int64":
		initialOrDefault := setValueToInitialOrDefault(f)
		u, err := strconv.ParseInt(initialOrDefault, 10, 64)
		if err != nil {
			log.Printf("Error converting value of %s to type int64\n", initialOrDefault)
			f.Type = ERROR_INCORRECT_TYPE
		}
		f.Value = u // int64
	}
}

func setFieldByName(c *Config, match string, field *Field) {
	for i, f := range c.Fields {
		if f.Name == match {
			*field = c.Fields[i]
		}
	}
}

func validate(r *http.Request, c *Config) {
	var fileErr error
	e := Error{
		Message: "",
		Type:    "",
	}
	if len(r.Form) == 0 {
		// Set all form fields as errored
		for i, _ := range c.Fields {
			c.Fields[i].Error.Type = ERROR_MISSING_VALUE
			setErrorMessage(&c.Fields[i], fileErr)
		}
		return
	}
	for key, value := range r.Form {
		val := strings.Join(value, "")
		for i, f := range c.Fields {
			if f.Name == key {
				c.Fields[i].Initial = val
				// Validate the field value
				if f.Validate {
					if val == "" || val == "<nil>" {
						e.Type = ERROR_MISSING_VALUE
					}
					if f.Type != "" {
						convertToType(&c.Fields[i])
						// Handle file validation
						if f.Type == "file" {
							cf, _, fileErr := r.FormFile(f.Name)
							if fileErr != nil {
								c.Fields[i].Error.Type = ERROR_FILE_TYPE
							}
							if cf == nil {
								c.Fields[i].Error.Type = ERROR_MISSING_VALUE
							}
						}
					} else {
						c.Fields[i].Value = val
					}
					// Test filed matches
				} else {
					// set the value for unvalidated fields
					c.Fields[i].Value = val
				}
				// Set Error Message
				c.Fields[i].Error = e
				// set the error message & pass in a fileErr which may or may not be nil
				setErrorMessage(&c.Fields[i], fileErr)
			}
		}
	}

	for i, f := range c.Fields {
		// If the form field undeclared then set an error
		if f.Validate && f.Value == nil {
			e.Type = ERROR_MISSING_VALUE
			c.Fields[i].Error = e
			setErrorMessage(&c.Fields[i], fileErr)
		}
		// All field values have been set on the config object - now perform matching validation
		if f.Matches != "" {
			var matchedField Field
			// Set a temporary matchedField var with the matching field only for matching
			setFieldByName(c, f.Matches, &matchedField)
			if f.Value != matchedField.Value {
				e.Type = ERROR_FIELDS_DO_NOT_MATCH
				c.Fields[i].Error = e
				setErrorMessage(&c.Fields[i], fileErr)
			}
		}
	}
}
