package icmp

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func Receive(timeChan chan time.Time) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		panic(err)
	}
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	receivedChan := make(chan bool)
	for {
		select {
		case startTime := <- timeChan:
			fmt.Printf("[pingo]: \tResponse Time: %s\n", time.Now().Sub(startTime).String())
		case  <- receivedChan:
			fmt.Printf("received!")
			//if received {
			//	fmt.Printf("[pingo]: \tResponse Time: %s\n", time.Now().Sub(startTime).String())
			//}
			break
		default:
			buf := make([]byte, 2048)
			_, err := f.Read(buf)
			if err != nil {
				log.Println(err)
			}

			receivedChan <- true
		}
	}

}


