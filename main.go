package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	stamp, err := getStamp()
	if err != nil {
		log.Fatal(err)
	}

	//Prod
	db, err := dbInit("/home/gnewton/gocode/src/github.com/gnewton/gohq/data/hq_slite3.db")

	// Test
	//db, err := dbInit("hq_slite3.db")
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
