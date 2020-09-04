package main

import (
	"fmt"
	"os"
    "github.com/dschoeffm/hetzner-dns-go"
)

func usage() {
    fmt.Println(os.Args[0], "<record type>", "<zone>", "<name>", "<value>",
        "<token>")
}

func main() {

    if len(os.Args) != 6 {
        usage()
        return
    }

    recordType := os.Args[1]
    zoneName := os.Args[2]
    name := os.Args[3]
    value := os.Args[4]
    token := os.Args[5]

	caller := hdns.GetApiCaller(token)
	zones, err := hdns.GetAllZones(caller)
	if err != nil {
		panic(err)
	}

	zoneId, err := hdns.IdOfZone(zoneName, zones)
	if err != nil {
		panic(err)
	}

	records, err := hdns.GetAllRecords(caller, zoneId)
	if err != nil {
		panic(err)
	}

	filteredRecords := hdns.FilterRecords(records, name, recordType)

	if len(filteredRecords) == 0 {
		// In this case, there are no records, create one
		fmt.Println("Creating record")
		err = hdns.CreateRecord(caller, &hdns.CreateRecordRequest{
			Type:   recordType,
			ZoneId: zoneId,
			Name:   name,
			Value:  value,
			Ttl:    1800})
		if err != nil {
			panic(err)
		}
	} else if len(filteredRecords) == 1 {
		// In this case, there is exatly one record
		fmt.Println("Updating record")
		rec := filteredRecords[0]
		if rec.Value != value {
			err = hdns.UpdateRecord(caller, &hdns.UpdateRecordRequest{
				Type:   recordType,
				ZoneId: zoneId,
				Name:   name,
				Value:  value,
				Ttl:    1800}, rec.Id)
			if err != nil {
				panic(err)
			}
		}
	} else {
		// Too many record... Don't know what to do... PANIC!!!!
		panic("More then one record exists...")
	}
}
