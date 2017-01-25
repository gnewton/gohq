package main

import (
	"fmt"
	"log"
)

func main() {
	// var file *os.File
	// var err error
	// path := "20170124103017.json"
	// if file, err = os.Open(path); err != nil {
	// 	log.Fatal(err)
	// }
	// reader := bufio.NewReader(file)
	// jsonStream, err := ioutil.ReadAll(reader)

	// reportStamp := "20170123164010"

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

	outages, err := getOutages(stamp)

	report, err := makeReport(outages, stamp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(report)

	persist(report, db)
}
