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
if ok, v := ValidateForm(r, Config{}); ok {
	// Form is valid
} else {
	// Handle form errors
}
```
Form with text / files
```go
if ok, v := ValidateMultiPartForm(r, Config{MaxMemory: 1200}); ok {
	// Form is valid
} else {
	// Handle form errors
}
```

