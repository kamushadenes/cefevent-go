package cefevent

import (
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type CEFExtension struct {
	ShortName   string
	FullName    string
	DataType    string
	Length      int
	Description string
}

type CEFEventExtension struct {
	CEFExtension *CEFExtension
	Value        interface{}
}

func GetEventExtension(name string) (*CEFEventExtension, error) {
	if v, ok := extensionShortMap[name]; ok {
		return &CEFEventExtension{v, nil}, nil
	}

	if v, ok := extensionLongMap[name]; ok {
		return &CEFEventExtension{v, nil}, nil
	}

	return nil, fmt.Errorf("invalid extension name %s", name)
}

func (e *CEFEventExtension) GetValueAsString() string {
	switch e.Value.(type) {
	case string:
		return e.Value.(string)
	case int64:
		return fmt.Sprintf("%d", e.Value.(int64))
	case float64:
		return fmt.Sprintf("%f", e.Value.(float64))
	default:
		return fmt.Sprintf("%v", e.Value)
	}
}

func (e *CEFEventExtension) SetString(v string) {
	v = strings.Replace(v, "\\", "\\\\", -1)
	v = strings.Replace(v, "=", "\\=", -1)
	v = strings.Replace(v, "\n", "\\n", -1)
	v = strings.TrimSpace(v)

	e.Value = v
}

func (e *CEFEventExtension) SetMACAddress(v string) {
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)

	e.Value = v
}

func (e *CEFEventExtension) SetInt(v int64) {
	e.Value = v
}

func (e *CEFEventExtension) SetLong(v float64) {
	e.Value = v
}

func (e *CEFEventExtension) SetFloatingPoint(v float64) {
	e.Value = v
}

func (e *CEFEventExtension) SetValue(v interface{}) error {
	switch e.CEFExtension.DataType {
	case "String", "TimeStamp", "IPv4 Address", "IPv6 Address":
		e.SetString(v.(string))
	case "MAC Address":
		e.SetMACAddress(v.(string))
	case "Integer":
		//e.SetInt(v.(int64))
		e.SetString(v.(string))
	case "Long":
		//e.SetLong(v.(float64))
		e.SetString(v.(string))
	case "Floating Point":
		//e.SetFloatingPoint(v.(float64))
		e.SetString(v.(string))
	default:
		return fmt.Errorf("invalid data type for extension %s: %s", e.CEFExtension.FullName, e.CEFExtension.DataType)
	}

	return nil
}

func (e *CEFEventExtension) Validate() (bool, error) {
	switch e.CEFExtension.DataType {
	case "String":
		switch t := e.Value.(type) {
		case string:
			v := e.Value.(string)
			if len(v) > e.CEFExtension.Length {
				return false, fmt.Errorf("invalid value for extension %s: max length is %d", e.CEFExtension.FullName, len(v))
			}
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "TimeStamp":
		// TODO Validate TimeStamp
	case "IPv4 Address":
		switch t := e.Value.(type) {
		case string:
			v := e.Value.(string)
			if !valid.IsIPv4(v) {
				return false, fmt.Errorf("invalid ipv4 for extension %s: %s", e.CEFExtension.FullName, v)
			}
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "IPv6 Address":
		switch t := e.Value.(type) {
		case string:
			v := e.Value.(string)
			if !valid.IsIPv6(v) {
				return false, fmt.Errorf("invalid ipv6 for extension %s: %s", e.CEFExtension.FullName, v)
			}
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "Integer":
		switch t := e.Value.(type) {
		case int64:
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "Long":
		switch t := e.Value.(type) {
		case float64:
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "Floating Point":
		switch t := e.Value.(type) {
		case float64:
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	case "MAC Address":
		switch t := e.Value.(type) {
		case string:
			v := e.Value.(string)
			if !valid.IsMAC(v) {
				return false, fmt.Errorf("invalid mac address for extension %s: %s", e.CEFExtension.FullName, v)
			}
		default:
			return false, fmt.Errorf("invalid value type for extension %s: %T", e.CEFExtension.FullName, t)
		}
	default:
		return false, fmt.Errorf("invalid data type for extension %s: %s", e.CEFExtension.FullName, e.CEFExtension.DataType)
	}

	return true, nil
}
