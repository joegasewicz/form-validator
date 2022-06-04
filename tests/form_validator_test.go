package tests

import (
	form_validator "github.com/joegasewicz/form-validator"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func createFormRequest(data url.Values, fn func(http.ResponseWriter, *http.Request)) {
	handler := fn
	// Create POST request
	r := httptest.NewRequest("POST", "/test", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	handler(w, r)
}

func TestValidFormContainsFieldMembers(t *testing.T) {
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

	// Create mock form
	data := url.Values{}
	data.Set("name", "Joe")
	data.Set("email", "joe@email.com")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
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
			t.Log("expected form to not error")
			t.Fail()
		}
	})
}

func TestFormFailsValueEmptyString(t *testing.T) {
	c := form_validator.Config{
		MaxMemory: 0,
		Fields: []form_validator.Field{
			{
				Name:     "name",
				Validate: true,
				Default:  "Joe",
				Type:     "string",
			},
		},
	}

	// Create mock form
	data := url.Values{}
	data.Set("name", "")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		if ok := form_validator.ValidateForm(r, &c); ok {
			t.Logf("expected form to fail validation\n")
			t.Fail()
		} else {
			name := c.Fields[0]
			if name.Value != "Joe" {
				t.Logf("expected Joe but got %s\n", name.Initial)
				t.Fail()
			}
		}
	})
}

func TestAllTypeConversionSuccessful(t *testing.T) {
	c := form_validator.Config{
		MaxMemory: 0,
		Fields: []form_validator.Field{
			{
				Name:     "name",
				Validate: true,
				Default:  "",
				Type:     "string",
			},
			{
				Name:     "is_bool",
				Validate: true,
				Default:  "",
				Type:     "bool",
			},
			{
				Name:     "is_float32",
				Validate: true,
				Default:  "",
				Type:     "float32",
			},
			{
				Name:     "is_float64",
				Validate: true,
				Default:  "",
				Type:     "float64",
			},
			{
				Name:     "is_uint8",
				Validate: true,
				Default:  "",
				Type:     "uint8",
			},
		},
	}

	// Create mock form
	data := url.Values{}
	data.Set("name", "Joe")
	data.Set("is_bool", "true")
	data.Set("is_float32", "0.12345679")
	data.Set("is_float64", "0.12345678912121212")
	data.Set("is_uint8", "255")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		if ok := form_validator.ValidateForm(r, &c); ok {

			name := c.Fields[0]
			is_bool := c.Fields[1]
			is_float32 := c.Fields[2]
			is_float64 := c.Fields[3]
			is_uint8 := c.Fields[4]
			if name.Value != "Joe" {
				t.Logf("expected Joe but got %s\n", name.Value)
				t.Fail()
			}
			if is_bool.Value != true {
				t.Logf("expected is_bool but got %s\n", is_bool.Value)
				t.Fail()
			}
			if is_float32.Value != float32(0.12345679) {
				t.Logf("expected is_float32 but got %s\n", is_float32.Value)
				t.Fail()
			}
			if is_float64.Value != 0.12345678912121212 {
				t.Logf("expected is_float64 but got %s\n", is_float64.Value)
				t.Fail()
			}
			if is_uint8.Value != uint8(255) {
				t.Logf("expected is_uint8 but got %s\n", is_uint8.Value)
				t.Fail()
			}
		} else {
			t.Logf("expected form to Pass validation\n")
			t.Fail()
		}
	})
}
