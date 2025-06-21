package form_validator

import (
	"github.com/stretchr/testify/assert"
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

func TestValidateMultiPartForm(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "name",
				Validate: true,
				Default:  "John",
				Type:     "string",
			},
			{
				Name:     "email",
				Validate: false,
				Type:     "string",
			},
		},
	}

	data := url.Values{}
	data.Set("name", "Joe")
	data.Set("email", "joe@email.com")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if ok := ValidateForm(r, &c); ok {
			name := c.Fields[0]
			email := c.Fields[1]
			if name.Name != "name" {
				t.Logf("expected name but got %s\n", name.Name)
				t.Fail()
			}
			if name.Value != "Joe" {
				t.Logf("expected Joe but got %s\n", name.Value)
				t.Fail()
			}
			if email.Value != "joe@email.com" {
				t.Logf("expected joe@email.com but got %s\n", email.Value)
				t.Fail()
			}

		}
	})
}

func TestValidFormContainsFieldMembers(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
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
		if ok := ValidateForm(r, &c); ok {
			// Name
			name := c.Fields[0]
			email := c.Fields[1]
			if name.Name != "name" {
				t.Logf("expected name but got %s\n", name.Name)
				t.Fail()
			}
			if name.Value != "Joe" {
				t.Logf("expected Joe but got %s\n", name.Value)
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
			if email.Value != "joe@email.com" {
				t.Logf("expected joe@email.com but got %s\n", email.Value)
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
				t.Logf("expected joe@email.com but got %s\n", email.Initial)
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
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
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
		if ok := ValidateForm(r, &c); ok {
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
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
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
			{
				Name:     "is_uint16",
				Validate: true,
				Default:  "",
				Type:     "uint16",
			},
			{
				Name:     "is_uint32",
				Validate: true,
				Default:  "",
				Type:     "uint32",
			},
			{
				Name:     "is_uint64",
				Validate: true,
				Default:  "",
				Type:     "uint64",
			},
			{
				Name:     "is_int8",
				Validate: true,
				Default:  "",
				Type:     "int8",
			},
			{
				Name:     "is_int16",
				Validate: true,
				Default:  "",
				Type:     "int16",
			},
			{
				Name:     "is_int32",
				Validate: true,
				Default:  "",
				Type:     "int32",
			},
			{
				Name:     "is_int64",
				Validate: true,
				Default:  "",
				Type:     "int64",
			},
			{
				Name:     "is_int",
				Validate: true,
				Default:  "",
				Type:     "int",
			},
			{
				Name:     "is_uint",
				Validate: true,
				Default:  "",
				Type:     "uint",
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
	data.Set("is_uint16", "65535")
	data.Set("is_uint32", "4294967295")
	data.Set("is_uint64", "18446744073709551615")
	data.Set("is_int8", "127")
	data.Set("is_int16", "32767")
	data.Set("is_int32", "2147483647")
	data.Set("is_int64", "18446744073709551615")
	data.Set("is_int", "100")
	data.Set("is_uint", "255")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		if ok := ValidateForm(r, &c); ok {

			name := c.Fields[0]
			is_bool := c.Fields[1]
			is_float32 := c.Fields[2]
			is_float64 := c.Fields[3]
			is_uint8 := c.Fields[4]
			is_uint16 := c.Fields[5]
			is_uint32 := c.Fields[6]
			is_uint64 := c.Fields[7]
			is_int8 := c.Fields[8]
			is_int16 := c.Fields[9]
			is_int32 := c.Fields[10]
			is_int64 := c.Fields[11]
			is_int := c.Fields[12]
			is_uint := c.Fields[13]

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
			if is_uint16.Value != uint16(65535) {
				t.Logf("expected is_uint16 but got %s\n", is_uint16.Value)
				t.Fail()
			}
			if is_uint32.Value != uint32(4294967295) {
				t.Logf("expected is_uint32 but got %s\n", is_uint32.Value)
				t.Fail()
			}
			if is_uint64.Value != uint64(18446744073709551615) {
				t.Logf("expected is_uint64 but got %s\n", is_uint64.Value)
				t.Fail()
			}
			if is_int8.Value != int8(127) {
				t.Logf("expected is_int8 but got %s\n", is_int8.Value)
				t.Fail()
			}
			if is_int16.Value != int16(32767) {
				t.Logf("expected is_int16 but got %s\n", is_int16.Value)
				t.Fail()
			}
			if is_int32.Value != int32(2147483647) {
				t.Logf("expected is_int32 but got %s\n", is_int32.Value)
				t.Fail()
			}
			if is_int64.Value != int64(9223372036854775807) {
				t.Logf("expected is_int64 but got %s\n", is_int64.Value)
				t.Fail()
			}
			if is_int.Value != 100 {
				t.Logf("expected is_int but got %s\n", is_int.Value)
				t.Fail()
			}
			if is_uint.Value != uint(255) {
				t.Logf("expected is_uint but got %s\n", is_uint.Value)
				t.Fail()
			}
		} else {
			t.Logf("expected form to Pass validation\n")
			t.Fail()
		}
	})
}

func TestGetFormGenericValue(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "weight",
				Validate: true,
				Default:  "",
				Type:     "float32",
			},
		},
	}

	// Create mock form
	data := url.Values{}
	data.Set("weight", "2.43")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		if ok := ValidateForm(r, &c); ok {
			weight := c.Fields[0]
			actual := getFormValue("weight", &c)
			if float32(2.43) != actual {
				t.Logf("expected 2.43 but got %s\n", weight.Initial)
				t.Fail()
			}
		} else {
			t.Logf("expected form to pass validation\n")
			t.Fail()
		}
	})
}

func TestGetFormError(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "weight",
				Validate: true,
				Default:  "",
				Type:     "float32",
			},
		},
	}

	// Create mock form
	data := url.Values{}
	data.Set("weight", "")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		if ok := ValidateForm(r, &c); ok {
			t.Logf("expected form to pass validation\n")
			t.Fail()
		} else {
			weight := c.Fields[0]
			actual := GetFormError("weight", &c)
			if actual.Type != ERROR_MISSING_VALUE {
				t.Logf("expected %s but got %s\n", ERROR_MISSING_VALUE, weight.Error.Type)
				t.Fail()
			}
			if len(actual.Message) == 0 {
				t.Logf("expected error message but got %s\n", actual.Message)
				t.Fail()
			}
		}
	})
}

func TestGetFormErrors(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		field   []Field
		wantRes map[string]map[string]string
	}{
		"exist error": {
			field: []Field{
				{
					Name: "weight",
					Error: Error{
						Message: "weight error message",
						Type:    "string",
					},
				},
				{
					Name: "height",
					Error: Error{
						Message: "height error message",
						Type:    "string",
					},
				},
			},
			wantRes: map[string]map[string]string{
				"weight": {
					"weight": "weight",
					"error":  "weight error message",
				}, "height": {
					"height": "height",
					"error":  "height error message",
				}},
		},
		"empty": {
			field:   []Field{},
			wantRes: map[string]map[string]string{},
		},
	}

	for name, tt := range testcases {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cfg := Config{
				MaxMemory: 0,
				Fields:    tt.field,
			}
			var formErrors = FormErrors{}
			GetFormErrors(&cfg, &formErrors)
			if len(tt.wantRes) == 0 {
				assert.Len(t, formErrors, 0, "formErrors must be empty")
			} else {
				for _, val := range tt.field {
					formError := formErrors[val.Name]
					assert.Equal(t, tt.wantRes[val.Name][val.Name], formError[val.Name], "name is not equal")
					assert.Equal(t, tt.wantRes[val.Name]["error"], formError["error"], "error message is not equal")
				}
			}
		})
	}
}

func TestMatchFields(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "password",
				Validate: true,
				Type:     "string",
			},
			{
				Name:     "confirm_password",
				Validate: true,
				Type:     "string",
				Matches:  "password",
			},
		},
	}

	// Create mock form
	data := url.Values{}
	data.Set("password", "wizard")
	data.Set("confirm_password", "blizzard")

	createFormRequest(data, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if ok := ValidateForm(r, &c); ok {
			t.Logf("expected form matches to fail")
			t.Fail()
		}
	})
}
