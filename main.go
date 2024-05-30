// Package main
package main

import (
	"log"
	"switcherctl/connection"
	"switcherctl/consts"
)

func main() {
	port, ok := consts.DeviceCategoryToUDPPort[consts.DeviceCategoryWaterHeater]
	if !ok {
		log.Fatalln("Could not find port for this device")
	}

	conn, err := connection.TryNew(consts.DefaultIP, port)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Fatalln(closeErr)
		}
	}()

	for {
		data, err := conn.Read()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Reply: %s\n", data)
	}
}
