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
	data.Set("email", "joe@email.com")
	// Create POST request
	r := httptest.NewRequest("POST", "/test", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	handler(w, r)
}

func TestValidateForm(t *testing.T) {
	c := form_validator.Config{
		MaxMemory: 0,
		Fields: []form_validator.Field{
			{
				Name:     "name",
				Validate: true,
				Default:  "John",
				Type:     "string",
			},
			{
				Name:     "email",
				Validate: false,
				Default:  "",
				Type:     "string",
			},
		},
	}

	createFormRequest(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if ok := form_validator.ValidateForm(r, &c); ok {
			// Name
			name := c.Fields[0]
			email := c.Fields[1]
			if name.Name != "name" {
				t.Logf("expected name but got %s\n", name.Name)
				t.Fail()
			}
			if name.Validate != true {
				t.Logf("expected true but got %v", name.Validate)
				t.Fail()
			}
			if name.Default != "John" {
				t.Logf("expected John but got %s\n", name.Default)
				t.Fail()
			}
			if name.Type != "string" {
				t.Logf("expected string but got %s\n", name.Type)
				t.Fail()
			}
			if name.Initial != "Joe" {
				t.Logf("expected Joe but got %s\n", name.Initial)
				t.Fail()
			}
			if name.Error.Message != "" {
				t.Logf("expected no errors but got %s\n", name.Error.Message)
			}
			// Email
			if email.Name != "email" {
				t.Logf("expected email but got %s\n", name.Name)
				t.Fail()
			}
			if email.Validate != false {
				t.Logf("expected false but got %v", email.Validate)
				t.Fail()
			}
			if email.Default != "" {
				t.Logf("expected '' but got %s\n", email.Default)
				t.Fail()
			}
			if email.Type != "string" {
				t.Logf("expected string but got %s\n", email.Type)
				t.Fail()
			}
			if email.Initial != "joe@email.com" {
				t.Logf("expected Joe but got %s\n", email.Initial)
				t.Fail()
			}
			if email.Error.Message != "" {
				t.Logf("expected no errors but got %s\n", email.Error.Message)
			}

		} else {
			t.Log("Should not fail")
			t.Fail()
		}
	})
}
