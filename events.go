package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

func findOutageEvent(outage *Outage, mostRecentReport *OutageReport, db *gorm.DB) int {
	if mostRecentReport == nil {
		return 0
	}
	for i, _ := range mostRecentReport.Outages {
		prevOutage := mostRecentReport.Outages[i]
		log.Println(prevOutage.ID, prevOutage.EventID)
		log.Println(outage.TimeStart, outage.Latitude, outage.Longitude)
		log.Println(prevOutage.TimeStart, prevOutage.Latitude, prevOutage.Longitude)

		if matches(outage, prevOutage) {
			log.Println("++++++++++++++++++++++++++")
			return prevOutage.EventID
		}
	}
	return getNewEventID(db)
}

func getNewEventID(db *gorm.DB) int {
	log.Println("")
	rows, err := db.Table("outages").Select("max(event_id)").Rows()
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("")
	if err != nil {
		log.Fatal(err)
		return 1
	}
	log.Println("")
	if rows == nil {
		return 1
	}

	var n int
	for rows.Next() {
		log.Println("")
		err = rows.Scan(&n)
		if err != nil {
			log.Println(err)
		}
		break
	}
	log.Println("")
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
