# Form Validator
Validate the incoming request's form values

# Install
```bash
go get github.com/joegasewicz/form-validator
```

### Setup
```go
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
```

# Example
Form with text fields
```go

if ok := ValidateForm(r, &c); ok {
	
} else {
	// Handle form errors
}
```
Form with files & text fields
```go
if ok := ValidateMultiPartForm(r, &c); ok {
	// Form is valid
} else {
	// Handle form errors
}
```
### Form Values
`GetFormValue` gets a single form field value

```go
name := GetFormValue("name", &c)
```

### Form Value Errors
`GetFormError` gets a single form error
```go
name := GetFormError("name", &c)
```

`GetAllFormErrors` gets all form errors
```go
var formErrs FormErrors = FormErrors{}
GetAllFormErrors(&c, &formErrs)
```

### Form Value Type Conversion
To convert a form value to a specific type, set the `Type` value in the `Field` struct, for example
```go
name := form_validator.Field{
    {
        Name:     "weight",
        Validate: true,
        Default:  "John",
        Type:     "float32",
    },
}
```
If the form successfully validates, the "weight" form value will be `float32(<VALUE>)`
The following type conversions are supported:
- string
- bool
- int, float32, float64
- int8, int16, int32, int64
- uint8, uint16, uint32, uint64
