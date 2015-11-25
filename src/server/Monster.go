package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"net/url"
)

func getRssFeed(keyword string, company string) []byte {
	
	requestUrl := fmt.Sprintf("http://rss.jobsearch.monster.com/rssquery.ashx?q=%s&cn=%s&rad_units=km&cy=DE&pp=2000&tm=14&sort=rv.di.dt&baseurl=stellenanzeige.monster.de", url.QueryEscape(keyword), url.QueryEscape(company))
	
	fmt.Println(requestUrl)
	resp, err := http.Get(requestUrl)

	var result []byte

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		
		result = contents
	}
	return result
}
