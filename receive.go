package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func recv() {
	fmt.Println("Starting socket!")
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		panic(err)
	}
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	for {
		buf := make([]byte, 2048)
		read, err := f.Read(buf)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("% X\n", buf[:read])
		fmt.Printf("%s \n", buf[28:read])
	}
}


