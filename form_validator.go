package form_validator

import (
	"fmt"
	"net/http"
)

// Config
type Config struct {
	MaxMemory int64
	Fields    *map[string]bool
}

// ValidateForm
func ValidateForm(r *http.Request, c Config) bool {
	validMetaData := make(map[string]string)
	r.ParseForm()
	validate(r, &validMetaData, c)
	return true
}

// ValidateMultiPartForm
func ValidateMultiPartForm(r *http.Request, c Config) bool {
	validMetaData := make(map[string]string)
	r.ParseMultipartForm(c.MaxMemory)
	validate(r, &validMetaData, c)
	return true
}

func validate(r *http.Request, v *map[string]string, c Config) {
	for key, value := range r.Form {
		if len(value) == 0 {
			fmt.Println("here-----> ", key)
		}
	}
}
