package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getStamp() (string, error) {
	url := "http://poweroutages.hydroquebec.com/pannes/donnees/v3_0/bisversion.json"
	resp, err := http.Get(url)
	//log.Println(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(body), "\""), nil

}

func getOutages(stamp string) ([]byte, *time.Time, error) {
	url := "http://poweroutages.hydroquebec.com/pannes/donnees/v3_0/bismarkers" + stamp + ".json"
	//log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	//log.Println(resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	ts := time.Now().UTC()
	return body, &ts, nil

}
