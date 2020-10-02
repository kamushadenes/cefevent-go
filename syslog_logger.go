package cefevent

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

type SyslogLogger struct {
	Protocol Protocol
	Host     string
	Port     int
	conn     net.Conn
}

func (l *SyslogLogger) Connect() error {
	switch l.Protocol {
	case TCP:
		return l.TCP()
	case UDP:
		return l.UDP()
	case LOCAL:
		return l.LOCAL()
	}

	return fmt.Errorf("unsupported protocol %s", l.Protocol)
}

func (l *SyslogLogger) IsConnected() bool {
	return l.conn != nil
}

func (l *SyslogLogger) TCP() error {
	if l.Protocol != TCP {
		return fmt.Errorf("invalid tcp protocol %s", l.Protocol)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", l.Host, l.Port))
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	l.conn = conn

	return nil
}

func (l *SyslogLogger) UDP() error {
	if l.Protocol != UDP {
		return fmt.Errorf("invalid udp protocol %s", l.Protocol)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", l.Host, l.Port))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}

	l.conn = conn

	return nil
}

func (l *SyslogLogger) LOCAL() error {
	return nil
}

func (l *SyslogLogger) Send(data []byte) error {
	switch l.Protocol {
	case TCP, UDP:
		_, err := l.conn.Write(data)
		return err
	default:
		return fmt.Errorf("unsupported protocol %s", l.Protocol)
	}
}

func (l *SyslogLogger) Close() error {
	switch l.Protocol {
	case TCP, UDP:
		return l.conn.Close()
	}

	return nil
}

func GetLogger(protocol Protocol, host string, port int) (*SyslogLogger, error) {
	var l SyslogLogger

	l.Protocol = protocol
	l.Host = host
	l.Port = port

	if err := l.Connect(); err != nil {
		return nil, err
	}

	return &l, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}
