package form_validator

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Config
type Config struct {
	MaxMemory int64
	Fields    []Field
}

type Field struct {
	Name     string
	Validate bool
	Default  string
	Type     string
	Value    interface{}
	Initial  string
	Error    Error
}

type Error struct {
	Message string
	Type    string
}

type jsonError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// ValidateForm
func ValidateForm(r *http.Request, c *Config) bool {
	r.ParseForm()
	validate(r, c)
	return true
}

// ValidateMultiPartForm
func ValidateMultiPartForm(r *http.Request, c *Config) bool {
	r.ParseMultipartForm(c.MaxMemory)
	validate(r, c)
	return true
}

func convertToType(f *Field) {
	switch f.Type {
	case "string":
		f.Value = f.Initial
		break
	case "float32":
		float, err := strconv.ParseFloat(f.Initial, 32)
		if err != nil {
			log.Printf("Error converting value of %s to type float32\n", f.Initial)
			f.Type = ERROR_INCORRECT_TYPE
		} else {
			f.Value = float
		}
		break
	case "float64":
		float, err := strconv.ParseFloat(f.Initial, 64)
		if err != nil {
			log.Printf("Error converting value of %s to type float64\n", f.Initial)
			f.Type = ERROR_INCORRECT_TYPE
		} else {
			f.Value = float
		}
		break
	case "bool":
	case "uint8":
	case "uint16":
	case "uint32":
	case "uint63":
	case "int8":
	case "int16":
	case "int32":
	case "int64":
	}
}

func validate(r *http.Request, c *Config) {
	for key, value := range r.Form {
		for _, f := range c.Fields {
			if f.Name == key {
				f.Initial = strings.Join(value, "")
				e := Error{
					Message: "",
					Type:    "",
				}
				if f.Validate {
					if len(value) == 0 {
						e.Type = ERROR_MISSING_VALUE
					}
					if f.Type != "" {
						convertToType(&f)
					}
				}
				// Set Error Message
				f.Error = e
				SetErrorMessage(&f)
			}
		}
	}
}
