# Form Validator
Validates incoming requests form values

# Install
```bash
go get github.com/joegasewicz/form-validator
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
