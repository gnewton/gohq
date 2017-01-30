package main

import (
	"log"
)

func main() {

	stamp, err := getStamp()
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbInit("/home/gnewton/gocode/src/github.com/gnewton/gohq/data/hq_slite3.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if containsStamp(stamp, db) {
		return
	}

	outages, timeAcquired, err := getOutages(stamp)

	report, err := makeReport(outages, stamp, timeAcquired)
	if err != nil {
		log.Println(string(outages))
		log.Fatal(err)
	}

	persist(report, db)
}
