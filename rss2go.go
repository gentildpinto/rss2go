package rss2go

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   uint16   `xml:"width"`
	Height  uint16   `xml:"height"`
}

type MediaContent struct {
	XMLName xml.Name `xml:"content"`
	Url     string   `xml:"url,attr"`
}

type Item struct {
	XMLName      xml.Name     `xml:"item"`
	Title        string       `xml:"title"`
	Link         string       `xml:"link"`
	Guid         string       `xml:"guid"`
	Description  string       `xml:"description"`
	MediaContent MediaContent `xml:"content"`
	Category     string       `xml:"category"`
	PubDate      string       `xml:"pubDate"`
}

type AtomLink struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	// Link        string   `xml:"link"` TODO: To solve -> It conflicts with AtomLink (<atom:link />)
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Copyright   string   `xml:"copyright"`
	AtomLink    AtomLink `xml:"http://www.w3.org/2005/Atom link"`
	Image       Image    `xml:"image"`
	Items       []Item   `xml:"item"`
}

func Rss2Go(feedUrl string) (*Feed, error) {
	var feed Feed

	resp, err := http.Get(feedUrl)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	re := regexp.MustCompile(`<!\[CDATA\[(.*?)\]\]>(\n*|\s*)`)
	newData := re.ReplaceAllString(string(data), "")
	data = []byte(newData)

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("something went wrong on unmarshal xml: %v", err)
	}

	return &feed, nil
}
