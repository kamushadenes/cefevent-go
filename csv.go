package cefevent

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCSV(filename string) ([]*CEFEvent, error) {
	var evts []*CEFEvent
	csvfile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvfile)

	var headers []string
	var cnt = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if cnt == 0 {
			headers = record
			cnt += 1
			continue
		}

		c := NewCEFEvent()

		for k := range headers {
			h := headers[k]

			switch h {
			case "name":
				c.SetName(record[k])
			case "deviceVendor":
				c.SetDeviceVendor(record[k])
			case "deviceProduct":
				c.SetDeviceProduct(record[k])
			case "signatureId":
				c.SetSignatureId(record[k])
			case "version":
				c.SetVersion(record[k])
			case "deviceVersion":
				c.SetDeviceVersion(record[k])
			case "severity":
				c.Severity = record[k]
			default:
				e, err := GetEventExtension(h)
				if err != nil {
					return nil, err
				}
				e.SetValue(record[k])

				c.Extensions = append(c.Extensions, e)
			}
		}
		evts = append(evts, c)
	}

	return evts, nil
}
