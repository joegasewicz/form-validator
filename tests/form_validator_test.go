package tests

import (
	"bytes"
	"fmt"
	form_validator "github.com/joegasewicz/form-validator"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkField(t *testing.T, err error) {
	if err != nil {
		t.Log("Failed to create mock form field")
		t.Fail()
	}
}

func TestValidateForm(t *testing.T) {
	f := make(map[string]bool)
	c := form_validator.Config{
		Fields: &f,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if ok := form_validator.ValidateForm(r, c); ok {
			fmt.Println("here-----> ")
		} else {
			t.Log("Should not fail")
			t.Fail()
		}
	}

	// Create mock form
	var form bytes.Buffer
	formWriter := multipart.NewWriter(&form)
	// name
	formWriter.WriteField("name", "joe")
	formWriter.Close()
	// Create POST request
	r := httptest.NewRequest("POST", "/test", &form)
	r.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	handler(w, r)

}
