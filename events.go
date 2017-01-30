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
	if &report == nil {
		return nil
	}
	populateReportOutages(&report, db)
	return nil
}

func populateReportOutages(report *OutageReport, db *gorm.DB) {
	var outages []*Outage
	db.Where("report_id = ?", report.Stamp).Find(&outages)
	for i, _ := range outages {
		log.Println(*outages[i])
	}
	report.Outages = outages
}
