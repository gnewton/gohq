package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

func findOutageEvent(outage *Outage, mostRecentReport *OutageReport, db *gorm.DB) {

}

func mostRecentReport(db *gorm.DB) *OutageReport {
	var report OutageReport
	db.Last(&report)
	//log.Println("000")
	//log.Println(report)
	//log.Println("111")
	if &report == nil {
		return nil
	}
	populateReportOutages(&report, db)
	return nil
}

func populateReportOutages(report *OutageReport, db *gorm.DB) {

}
