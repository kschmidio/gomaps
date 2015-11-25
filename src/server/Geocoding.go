package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var geocodingUrl string = "https://maps.googleapis.com/maps/api/geocode/json?address="

func geocode(cityName string) []byte {
	
	request := geocodingUrl + cityName
	resp, err := http.Get(request)

	var result []byte

	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			  os.Exit(1)
		}
		
		result = contents
	} else {
		fmt.Printf("%s", err)
	}
	return result
}
