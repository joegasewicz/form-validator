# Form Validator
Validate the incoming request's form values

# Install
```bash
go get github.com/joegasewicz/form-validator
```

### Setup
```go
c := Config{
    MaxMemory: 0,
    Fields: []Field{
        {
            Validate: true,
            Default:  "John",
            Type:     "string",
        },
        {
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

if ok, json := ValidateForm(r, &c); ok {
	// Form is valid
} else {
	// Handle form errors
}
```
Form with text / files
```go
if ok, json := ValidateMultiPartForm(r, &c); ok {
	// Form is valid
} else {
	// Handle form errors
}
```

