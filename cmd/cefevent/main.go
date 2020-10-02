package main

import (
	"cefevent"
	"fmt"
)

func main() {
	m := cefevent.NewLoggerManager()
	m.Start(120000)

	l, err := cefevent.GetLogger("udp", "localhost", 9999)
	if err != nil {
		panic(err)
	}

	m.AddLogger("test", l)

	//m.AutoSend([][]byte{[]byte("test")})

	evts, err := cefevent.ReadCSV("sample_logs.csv")
	fmt.Println(evts, err)
}
