package main

import (
    "bytes"
    "encoding/xml"
    "fmt"
    "os"
)

type RssFeed struct {
    XMLName xml.Name    `xml:"rss"`
    Channel *RssChannel `xml:"channel"`
}

type RssChannel struct {
    XMLName       xml.Name `xml:"channel"`
    Title         string   `xml:"title"`
    Description   string   `xml:"description"`
    Link          string   `xml:"link"`
    Language      string   `xml:"language"`
    PubDate       string   `xml:"pubDate"`
    LastBuildDate string   `xml:"lastBuildDate"`
    ItemList	  []Item   `xml:"item"`
}

type Item struct {
	Title		string		`xml:"title"`
	Description	string		`xml:"description"`
	Link		string		`xml:"link"`	
	PubDate		string		`xml:"pubDate"`
}

func parseXml(x []byte) *RssFeed {
    feed := RssFeed{}
    d := xml.NewDecoder(bytes.NewReader(x))
    err := d.Decode(&feed)
    if err != nil {
        fmt.Println("Failed decoding xml")
        os.Exit(0)
    }
    return &feed
}

