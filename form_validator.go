package form_validator

import (
	"fmt"
	"net/http"
)

type Config struct {
	MaxMemory int64
}

func ValidateForm(r *http.Request, c Config) {
	validMetaData := make(map[string]string)
	r.ParseForm()
	validate(r, &validMetaData, c)
}

func ValidateMultiPartForm(r *http.Request, c Config) {
	validMetaData := make(map[string]string)
	r.ParseMultipartForm(c.MaxMemory)
	validate(r, &validMetaData, c)
}

func validate(r *http.Request, v *map[string]string, c Config) {
	for key, value := range r.Form {
		if len(value) == 0 {
			fmt.Println("here-----> ", key)
		}
	}
}
