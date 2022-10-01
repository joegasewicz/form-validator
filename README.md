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
            Name:     "email",
            Validate: true,
            Type:     "string",
        },
        {
            Name:     "email",
            Validate: true,
            Type:     "string",
        },
        {
            Name:     "confirm_email",
            Validate: true,
            Type:     "string",
            Matches:  "email", // checks that email & confirm_email match
        },
    },
}

if ok := form_validator.ValidateForm(r, &c); ok {
    eail, _ := form_validator.GetString("email", &c)
    password, _ := form_validator.GetString("password", &c)
} else {
	// Get all the form errors
    var formErrs = form_validator.FormErrors{}
    form_validator.GetFormErrors(&c, &formErrs) 
}

```
The Field type properties are:
 - Name field is the form's 'name' value
 - Validate sets whether the field requires validation
 - Default set a default value is the form field empty
 - Type sets the type conversion e.g. int8, uint, float16 ...
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

### Match field values (password confirmation)
If you require password fields, for example to be matched, then assign a `Matches` value to a field:
```go
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
```
The above validation will fail if the `password` field's value is not the same as the `confirm_password` field.

### Form Value Errors
`GetFormError` gets a single form error
```go
name := GetFormError("name", &c)
```
### GetFormErrors
`GetFormErrors` access all form errors as a map (`FormErrors`) indexed off the form names
```go
var formErrs = form_validator.FormErrors{}
form_validator.GetFormErrors(&c, &formErrs)
```
If the results of `formErrs` are passed to the template as data then
all form errors can be accessed from the map via index name, for example:
```go
{{ if .FormErrors.title }}
           <div class="alert alert-danger" role="alert">
               {{ .FormErrors.title.error }}
           </div>
 {{ end }}
```
In this case `FormErrors.title.error` will produce an error message that
can be safely displayed to the user.

#### Get the form field value's correct value & type
There are `Get<TYPE>(name string, *Config)` functions for each supported type.
For example
```go
var title string
title, _ = GetString("title", &c)

var id int32
id, _ = GetInt32("id", &c)
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
- file
- int, float32, float64
- int8, int16, int32, int64
- uint8, uint16, uint32, uint64

