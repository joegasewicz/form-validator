package form_validator

import (
	"fmt"
	"strconv"
)

func getFormValue(name string, c *Config) interface{} {
	for _, v := range c.Fields {
		if v.Name == name {
			return v.Value
		}
	}
	return nil
}

// GetString gets any string types from the form values
//
//
// 		myStr, _ = GetString("name", &c)
//
func GetString(name string, c *Config) (string, error) {
	return fmt.Sprintf("%v", getFormValue(name, c)), nil
}

// GetBool gets any bool types from the form values
//
//
// 		myBool, _ = GetBool("is_happy", &c)
//
func GetBool(name string, c *Config) (bool, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	if b == "true" || b == "True" {
		return true, nil
	}
	i, err := strconv.Atoi(b)
	if i == 1 {
		return true, err
	}
	return false, err
}

// GetFloat32 gets any int types from the form values
//
//
// 		myFloat32, _ = GetFloat32("age", &c)
//
func GetFloat32(name string, c *Config) (float32, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	float, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return 0, err
	}
	return float32(float), err
}

// GetFloat64 gets any int types from the form values
//
//
// 		myFloat64, _ = GetFloat64("age", &c)
//
func GetFloat64(name string, c *Config) (float64, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	float, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, err
	}
	return float, err
}

// GetInt gets any int types from the form values
//
//
// 		myInt, _ = GetInt("age", &c)
//
func GetInt(name string, c *Config) (int, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	i, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return i, err
}

// GetInt8 gets any int types from the form values
//
//
// 		myInt8, _ = GetInt8("age", &c)
//
func GetInt8(name string, c *Config) (int8, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseInt(b, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(u), err
}

// GetInt16 gets any int types from the form values
//
//
// 		myInt16, _ = GetInt16("age", &c)
//
func GetInt16(name string, c *Config) (int16, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseInt(b, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(u), err
}

// GetInt32 gets any int types from the form values
//
//
// 		myInt32, _ = GetInt32("age", &c)
//
func GetInt32(name string, c *Config) (int32, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseInt(b, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(u), err
}

// GetInt64 gets any int types from the form values
//
//
// 		myInt64, _ = GetInt64("age", &c)
//
func GetInt64(name string, c *Config) (int64, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(u), err
}

// GetUint gets any uint types from the form values
//
//
// 		GetUint, _ = GetUint("age", &c)
//
func GetUint(name string, c *Config) (uint, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseUint(b, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u), err
}

// GetUint8 gets any uint types from the form values
//
//
// 		GetUint8, _ = GetUint8("age", &c)
//
func GetUint8(name string, c *Config) (uint8, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseUint(b, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(u), err
}

// GetUint16 gets any uint types from the form values
//
//
// 		GetUint16, _ = GetUint16("age", &c)
//
func GetUint16(name string, c *Config) (uint16, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseUint(b, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(u), err
}

// GetUint32 gets any uint types from the form values
//
//
// 		GetUint32, _ = GetUint32("age", &c)
//
func GetUint32(name string, c *Config) (uint32, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseUint(b, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), err
}

// GetUint64 gets any uint types from the form values
//
//
// 		GetUint64, _ = GetUint64("age", &c)
//
func GetUint64(name string, c *Config) (uint64, error) {
	var err error
	b := fmt.Sprintf("%v", getFormValue(name, c))
	u, err := strconv.ParseUint(b, 10, 32)
	if err != nil {
		return 0, err
	}
	return u, err
}
