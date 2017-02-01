package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

func findOutageEvent(outage *Outage, mostRecentReport *OutageReport, db *gorm.DB) {
	if mostRecentReport == nil {
		return
	}
	for i, _ := range mostRecentReport.Outages {
		prevOutage := mostRecentReport.Outages[i]
		log.Println("")
		log.Println(outage.TimeStart, outage.Latitude, outage.Longitude)
		log.Println(prevOutage.TimeStart, prevOutage.Latitude, prevOutage.Longitude)

		if matches(outage, prevOutage) {
			log.Println("++++++++++++++++++++++++++")
			outage.EventID = prevOutage.EventID
			return
		}
	}
	outage.EventID = getNewEventID(db)
}

func getNewEventID(db *gorm.DB) uint {

	rows, err := db.Table("outages").Select("max(event_id)").Rows()

	if err != nil {
		log.Fatal(err)
		return 1
	}

	if rows == nil {
		return 1
	}

	var n uint
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			log.Println(err)
		}
		break
	}
	return n + 1
}

func matches(n, o *Outage) bool {
	return /*n.TimeStart == o.TimeStart && */ n.Latitude == o.Latitude && n.Longitude == o.Longitude && n.ClientsEffected == o.ClientsEffected
}

func mostRecentReport(db *gorm.DB) *OutageReport {
	var count int
	err := db.Table("outage_reports").Count(&count)
	if err != nil {
		//log.Printf("%+v\n", err)
		//log.Println(err)
	}
	log.Println("count=", count)
	if count == 0 {
		return nil
	}

	var report OutageReport
	err = db.Last(&report)
	//if err != nil {
	//log.Fatal(err)
	//}
	if &report == nil {
		return nil
	}
	populateReportOutages(&report, db)
	return &report
}

func populateReportOutages(report *OutageReport, db *gorm.DB) {
	var outages []*Outage
	db.Where("report_id = ?", report.Stamp).Find(&outages)

	report.Outages = outages
}
