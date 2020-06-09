package main

import "flag"

func main() {
	destination := *flag.String("dest", "127.0.0.1", "The IP address of the destination server. " +
		"Note this should not contain a port number. Examples: '127.0.0.1', '8.8.8.8'")

}