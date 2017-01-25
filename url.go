package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getStamp() (string, error) {
	url := "http://poweroutages.hydroquebec.com/pannes/donnees/v3_0/bisversion.json"
	resp, err := http.Get(url)
	log.Println(url)
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

func getOutages(stamp string) ([]byte, error) {
	url := "http://poweroutages.hydroquebec.com/pannes/donnees/v3_0/bismarkers" + stamp + ".json"
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
