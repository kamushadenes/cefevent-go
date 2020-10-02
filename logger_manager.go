package cefevent

import (
	"math/rand"
	"time"

	"github.com/alitto/pond"
)

type LoggerManager struct {
	Loggers map[string]Logger
	eps     int
	pool    *pond.WorkerPool
}

func NewLoggerManager() *LoggerManager {
	return &LoggerManager{}
}

func (m *LoggerManager) Start(eps int) {
	m.eps = eps

	m.pool = pond.New(m.eps, 0, pond.Strategy(pond.Lazy()), pond.MinWorkers(m.eps))
}

func (m *LoggerManager) StopAndWait() {
	m.pool.StopAndWait()
}

func (m *LoggerManager) AddLogger(name string, l Logger) error {
	if m.Loggers == nil {
		m.Loggers = make(map[string]Logger)
	}

	if !l.IsConnected() {
		if err := l.Connect(); err != nil {
			return err
		}
	}

	m.Loggers[name] = l

	return nil
}

func (m *LoggerManager) RemoveLogger(name string) {
	if m.Loggers == nil {
		m.Loggers = make(map[string]Logger)
	}

	if _, ok := m.Loggers[name]; ok {
		delete(m.Loggers, name)
	}
}

func (m *LoggerManager) Send(msg []byte) {
	for k := range m.Loggers {
		m.pool.Submit(func() {
			m.Loggers[k].Send(msg)
		})
	}
}

func (m *LoggerManager) AutoSend(msgs [][]byte) {
	ticker := time.NewTicker(time.Duration(1000000000/(m.eps/10*35)) * time.Nanosecond)
	for {
		m.Send(msgs[rand.Intn(len(msgs))])
		<-ticker.C
	}
}
