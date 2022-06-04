package tests

import (
	form_validator "github.com/joegasewicz/form-validator"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func checkField(t *testing.T, err error) {
	if err != nil {
		t.Log("Failed to create mock form field")
		t.Fail()
	}
}

func createFormRequest(fn func(http.ResponseWriter, *http.Request)) {
	handler := fn
	// Create mock form
	data := url.Values{}
	data.Set("name", "Joe")
	// Create POST request
	r := httptest.NewRequest("POST", "/test", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	handler(w, r)
}

func TestValidateForm(t *testing.T) {
	f := make(map[string]interface{})
	f["name"] = map[string]interface{}{
		"validate": true,
		"default":  "default name",
		"type":     "string",
	}
	c := form_validator.Config{
		Fields: &f,
	}

	createFormRequest(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if ok := form_validator.ValidateForm(r, c); ok {
			// Passed
		} else {
			t.Log("Should not fail")
			t.Fail()
		}
	})
}
