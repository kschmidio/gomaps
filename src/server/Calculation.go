package main

import (
	"strings"
	"fmt"
	"time"
)

func calculateJobsPerCity(rss *RssChannel) []JobDTO {
	m := make(map[string]int)

	var jobDTOs []JobDTO
	
	for _, element := range rss.ItemList {
		cities := strings.Split(element.Description, ",")
		city := strings.Split(cities[0], " ")[0]
		
		m[city] += 1
		pubDate, _ := time.Parse(time.RFC1123, element.PubDate)
		now := time.Now()
		
		freshPub := pubDate.YearDay() >= now.YearDay() - 1
		jobDTO := JobDTO{CityName : city, JobDescription : element.Description, Title : element.Title, Link : element.Link, Today : freshPub}
		jobDTOs = append(jobDTOs, jobDTO) 
	}

	cityCoords := readFromFile()
	
	for key, _ := range m {
		_ , ok := cityCoords[key]
		if (!ok || cityCoords[key].Lat < 1) {
			
			fmt.Println("geocoding city ", key)
			/*
			geocodingJson := geocode(key)
			geocodingResponse := parseJson(geocodingJson)
			
			if (len(geocodingResponse.Result) > 0) {
				cityCoords[key] = geocodingResponse.Result[0].Geo.Loc		
			} else {
				fmt.Println("no coordinates found for ", geocodingResponse)			
			}
			*/
		}
	}
	//saveToFile(cityCoords)
	
	for index, element := range jobDTOs {
		lat := 0.0001 * float64(index)
		lng := 0.0001 * float64(index)
		jobDTOs[index].Lat = cityCoords[element.CityName].Lat + lat
		jobDTOs[index].Lng = cityCoords[element.CityName].Lng + lng
	}	
	
	return jobDTOs
}
