package cefevent

import (
	"errors"
	"time"

	"github.com/influxdata/go-syslog/v2/rfc5424"
)

func BuildSyslogMessage(timestamp *time.Time, facility Facility, severity Severity, message string) (*rfc5424.SyslogMessage, error) {
	return buildSyslogMessage(timestamp, facility, severity, "", "", "", "", message)
}

func buildSyslogMessage(timestamp *time.Time, facility Facility, severity Severity, hostname Hostname, appId string, procId string, msgId string, message string) (*rfc5424.SyslogMessage, error) {
	msg := &rfc5424.SyslogMessage{}

	msg.SetVersion(1)

	msg.SetTimestamp(timestamp.Format(time.RFC3339))

	msg.SetPriority(uint8(severity) | uint8(facility))

	if hostname != "" {
		msg.SetHostname(string(hostname))
	}

	if appId != "" {
		msg.SetAppname(appId)
	}

	if procId != "" {
		msg.SetProcID(procId)
	}

	if msgId != "" {
		msg.SetMsgID(msgId)
	}

	msg.SetMessage(message)

	if valid := msg.Valid(); valid {
		return msg, nil
	}

	return msg, errors.New("invalid message")
}
