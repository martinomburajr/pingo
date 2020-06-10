package main

import (
	"flag"
	"fmt"
	"github.com/martinomburajr/pingo/icmp"
	"math"
	"time"
)

func main() {
	var argDestinationIP string
	var argPingCount int

	flag.StringVar(&argDestinationIP, "dest", "127.0.0.1",
		"The IP address of the destinationIP server. " +
		"Note this should not contain a port number. Examples: '127.0.0.1', '8.8.8.8'")
	flag.IntVar(&argPingCount,"count", 5, "The number of times to Send a ping request.")
	flag.Parse()

	if argPingCount == -1 {
		argPingCount = math.MaxInt32
	}

	timeChan := make(chan time.Time)
	fmt.Println(argDestinationIP)

	go icmp.Receive(timeChan)
	for i := 0; i < argPingCount; i++ {
		icmp.Send(argDestinationIP, timeChan)
		time.Sleep(1000 * time.Millisecond)
	}
}
