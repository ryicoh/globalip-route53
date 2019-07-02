package main

import (
	"log"
)

func main() {
	log.Println("Zone Id: " + getZoneId())
	log.Println("Record Name: " + getRecordName())
	globalIp := getGlobalIp()
	log.Println("Global Ip: " + globalIp)

	recordIp := getRecordIp()
	if recordIp != globalIp {
		changeRecord(globalIp)
	}
}
