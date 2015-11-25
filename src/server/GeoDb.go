package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"io"
	//"fmt"
)

func saveToFile(locOfCities map[string]Location) {
	f, _ := os.Create("geodb")
	
	w := csv.NewWriter(f)
	
	for key, value := range locOfCities {
		csvRecord := []string{key, strconv.FormatFloat(value.Lat, 'f', -1, 64), strconv.FormatFloat(value.Lng, 'f', -1, 64)}
		if err := w.Write(csvRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func readFromFile() map[string]Location {
	f, _ := os.Open("geodb")
	r := csv.NewReader(f)
	
	cityCoords := make(map[string]Location)
	
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		lat, _ := strconv.ParseFloat(record[1], 64)
		lng, _ := strconv.ParseFloat(record[2], 64)
		cityCoords[record[0]] = Location{lat, lng}
		//fmt.Println("city ", record[0], " lat " , record[1], " lng ", record[2])
	}
	
	return cityCoords
}