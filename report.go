package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

const stampForm = "20060102150405 MST"
const longForm = "2006-01-02 15:04:05 MST"
const EST = " EST"

type OutageReport struct {
	Stamp        string `sql:"size:14" gorm:"primary_key"`
	TimeAcquired *time.Time
	Outages      []*Outage `sql:"-"`
	JsonGz       []byte
}

type Outage struct {
	ID       uint   `gorm:"primary_key"`
	EventID  uint   `gorm:"index"`
	ReportID string `sql:"size:14;index" `
	//ReportStamp     time.Time
	ClientsEffected int
	TimeStart       time.Time
	TimeEndEstimate time.Time
	EventCode       string `sql:"size:1"`
	Latitude        float64
	Longitude       float64
	OtherCode       string `sql:"size:1"`
	Other1          string `sql:"size:32"`
	Other2          string `sql:"size:32"`
	Other3          string `sql:"size:32"`
	Other4          string `sql:"size:32"`
}

func makeReport(jsonStream []byte, reportStamp string, timeAcquired *time.Time) (*OutageReport, error) {

	var o OutagesJson

	err := json.Unmarshal(jsonStream, &o)
	if err != nil {
		log.Println("foo", err)
		return nil, err
	}
	report := new(OutageReport)
	report.Stamp = reportStamp
	report.TimeAcquired = timeAcquired

	for i, _ := range o.Outages {
		newOutage, err := makeOutage(o.Outages[i], i)
		if err != nil {

			return nil, err
		}
		report.Outages = append(report.Outages, newOutage)

		newOutage.ReportID = reportStamp
		//newOutage.ReportStamp, err = time.Parse(stampForm, reportStamp+EST)
		if err != nil {
			return nil, err
		}

	}
	return report, nil
}

func makeOutage(d []interface{}, orderInReport int) (*Outage, error) {

	outage := new(Outage)

	if clientsEffected, ok := d[ClientsEffected].(float64); ok {
		outage.ClientsEffected = int(clientsEffected)
	} else {
		return nil, errors.New("Num clients not float64")
	}

	if timeStart, ok := d[TimeStart].(string); ok {
		var err error
		outage.TimeStart, err = time.Parse(longForm, timeStart+EST)
		outage.TimeStart = outage.TimeStart.UTC()
		if err != nil {
			log.Println(timeStart)
			log.Println(err)
		}
	}

	if timeEndEstimate, ok := d[TimeEndEstimate].(string); ok {
		if timeEndEstimate != "" {
			var err error
			outage.TimeEndEstimate, err = time.Parse(longForm, timeEndEstimate+EST)
			outage.TimeEndEstimate = outage.TimeEndEstimate.UTC()
			if err != nil {
				log.Println("444")
				log.Println(timeEndEstimate)
				log.Println(err)
			}
		}
	}

	if eventCode, ok := d[EventCode].(string); ok {
		outage.EventCode = eventCode
	}

	if latlong, ok := d[LatLong].(string); ok {
		latlong = string(latlong[1 : len(latlong)-1])
		parts := strings.Split(latlong, ", ")
		if parts == nil || len(parts) != 2 {
			return nil, errors.New("Incorrect lat/long structure")
		}
		var err error
		outage.Latitude, err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}
		outage.Longitude, err = strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return nil, err
		}
	}

	if otherCode, ok := d[OtherCode].(string); ok {
		outage.OtherCode = otherCode
	}

	if other1, ok := d[Other1].(string); ok {
		outage.Other1 = other1
	}

	if other2, ok := d[Other2].(string); ok {
		outage.Other2 = other2
	}

	if other3, ok := d[Other3].(string); ok {
		outage.Other3 = other3
	}

	if other4, ok := d[Other4].(string); ok {
		outage.Other4 = other4
	}

	return outage, nil
}
