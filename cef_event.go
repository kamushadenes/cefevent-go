package cefevent

import (
	"fmt"
	"strings"
)

type CEFEvent struct {
	Name          string
	DeviceVendor  string
	DeviceProduct string
	SignatureId   string
	Version       string
	DeviceVersion string
	Severity      string
	Extensions    []*CEFEventExtension
}

func NewCEFEvent() *CEFEvent {
	var c CEFEvent

	c.Reset()

	return &c
}

func (c *CEFEvent) Reset() {
	c.Version = "0"
	c.DeviceVendor = "CEF Vendor"
	c.DeviceProduct = "CEF Product"
	c.DeviceVersion = "1.0"
	c.SignatureId = "0"
	c.Name = "CEF Event"
	c.Severity = "Medium"

	c.Extensions = nil
}

func (c *CEFEvent) String() string {
	msg := fmt.Sprintf("CEF:%s|%s|%s|%s|%s|%s|%s|", c.Version, c.DeviceVendor, c.DeviceProduct, c.DeviceVersion, c.SignatureId, c.Name, c.Severity)

	for k := range c.Extensions {
		e := c.Extensions[k]
		msg += fmt.Sprintf("%s=%s ", e.CEFExtension.ShortName, e.GetValueAsString())
	}

	msg = strings.TrimSpace(msg)

	return msg
}

func (c *CEFEvent) Bytes() []byte {
	return []byte(c.String())
}

func (c *CEFEvent) formatPrefix(p string) string {
	p = strings.Replace(p, "\\", "\\\\", -1)
	p = strings.Replace(p, "|", "\\|", -1)
	p = strings.TrimSpace(p)

	return p
}

func (c *CEFEvent) SetName(v string) error {
	c.Name = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetDeviceVendor(v string) error {
	c.DeviceVendor = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetDeviceProduct(v string) error {
	c.DeviceProduct = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetSignatureId(v string) error {
	c.SignatureId = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetVersion(v string) error {
	c.Version = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetDeviceVersion(v string) error {
	c.DeviceVersion = c.formatPrefix(v)

	return nil
}

func (c *CEFEvent) SetSeverityString(v string) error {
	switch v {
	case "Unknown", "Low", "Medium", "High", "Very-High":
		c.Severity = v
	default:
		return fmt.Errorf("invalid severity string %s", v)
	}

	return nil
}

func (c *CEFEvent) SetSeverityIntMappedToString(v int) error {
	switch v {
	case 0, 1, 2, 3:
		c.Severity = "Low"
	case 4, 5, 6:
		c.Severity = "Medium"
	case 7, 8:
		c.Severity = "High"
	case 9, 10:
		c.Severity = "Very-High"
	default:
		return fmt.Errorf("invalid severity int %d", v)
	}

	return nil
}

func (c *CEFEvent) SetSeverityInt(v int) error {
	if v < 0 || v > 10 {
		return fmt.Errorf("invalid severity int %d", v)
	}

	c.Severity = fmt.Sprintf("%d", v)

	return nil
}
