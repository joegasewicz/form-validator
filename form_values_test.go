package form_validator

import "testing"

func TestGetString(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "name",
				Validate: true,
				Default:  "John",
				Type:     "string",
				Value:    "Joe",
			},
		},
	}

	var expected, actual string
	expected = "Joe"
	actual, _ = GetString("name", &c)
	if expected != actual {
		t.Logf("expected %s but got %s", expected, actual)
	}
}

func TestGetBool(t *testing.T) {
	c := Config{
		MaxMemory: 0,
		Fields: []Field{
			{
				Name:     "name1",
				Validate: true,
				Default:  "John",
				Type:     "bool",
				Value:    "true",
			},
			{
				Name:     "name2",
				Validate: true,
				Default:  "John",
				Type:     "bool",
				Value:    "true_",
			},
			{
				Name:     "name3",
				Validate: true,
				Default:  "John",
				Type:     "bool",
				Value:    "True",
			},
			{
				Name:     "name4",
				Validate: true,
				Default:  "John",
				Type:     "bool",
				Value:    1,
			},
		},
	}

	var expected, actual bool

	expected = true
	actual, _ = GetBool("name1", &c)
	if actual != expected {
		t.Logf("Expected %v but got %v", expected, actual)
		t.Fail()
	}

	expected = false
	actual, _ = GetBool("name2", &c)
	if actual != expected {
		t.Logf("Expected %v but got %v", expected, actual)
		t.Fail()
	}

	expected = true
	actual, _ = GetBool("name3", &c)
	if actual != expected {
		t.Logf("Expected %v but got %v", expected, actual)
		t.Fail()
	}

	expected = true
	actual, _ = GetBool("name4", &c)
	if actual != expected {
		t.Logf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}
