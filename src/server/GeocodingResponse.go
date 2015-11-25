package main

import (
	"encoding/json"
	"fmt"
)

type GeocodingResponse struct {
	Result []Results `json:"results"`
	StatusCode 	string `json:"status"`
}

type Results struct {
	Geo	Geometry `json:"geometry"`
}

type Geometry struct {
	Loc Location	`json:"location"`
}

type Location struct {
	Lat	float64	`json:"lat"`
	Lng	float64	`json:"lng"`
}

func parseJson(jsonBlob []byte) GeocodingResponse {
    var response GeocodingResponse

	err := json.Unmarshal(jsonBlob, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	
    return response
}

