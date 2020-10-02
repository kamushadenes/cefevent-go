package cefevent

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

type Severity uint8
type Facility uint8

type Protocol string

const (
	EMERGENCY Severity = 0
	ALERT     Severity = 1
	CRITICAL  Severity = 2
	ERROR     Severity = 3
	WARNING   Severity = 4
	NOTICE    Severity = 5
	INFO      Severity = 6
	DEBUG     Severity = 7

	KERN     Facility = 0
	USER     Facility = 1
	MAIL     Facility = 2
	DAEMON   Facility = 3
	AUTH     Facility = 4
	SYSLOG   Facility = 5
	LPR      Facility = 6
	NEWS     Facility = 7
	UUCP     Facility = 8
	CRON     Facility = 9
	AUTHPRIV Facility = 10
	FTP      Facility = 11
	NTP      Facility = 12
	SECURITY Facility = 13
	CONSOLE  Facility = 14
	CLOCK    Facility = 15
	LOCAL0   Facility = 16
	LOCAL1   Facility = 17
	LOCAL2   Facility = 18
	LOCAL3   Facility = 19
	LOCAL4   Facility = 20
	LOCAL5   Facility = 21
	LOCAL6   Facility = 22
	LOCAL7   Facility = 23

	TCP   Protocol = "tcp"
	UDP   Protocol = "udp"
	LOCAL Protocol = "local"
)

type Hostname string

func NewHostname(hostname string) (Hostname, error) {
	if !valid.IsHost(hostname) {
		return "", fmt.Errorf("invalid hostname %s", hostname)
	}

	return Hostname(hostname), nil
}
