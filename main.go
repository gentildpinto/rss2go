package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
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

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Guid        string   `xml:"guid"`
	Description string   `xml:"description"`
}

type AtomLink struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Copyright   string   `xml:"copyright"`
	Image       Image    `xml:"image"`
	AtomLink    AtomLink `xml:"http://www.w3.org/2005/Atom link"`
	Items       []Item   `xml:"item"`
}

func main() {
	handler := http.HandlerFunc(getXml)
	http.Handle("/", handler)
	http.ListenAndServe(":8089", nil)
}

func getXml(w http.ResponseWriter, r *http.Request) {
	data, _ := GetXML("https://g1.globo.com/rss/g1")
	// fmt.Println(string(data))

	re := regexp.MustCompile(`<!\[CDATA\[(.*?)\]\]>(\n*|\s*)`)
	newData := re.ReplaceAllString(string(data), "")
	data = []byte(newData)
	// for index, item := range feed.Channel.Items {
	// 	feed.Channel.Items[index].Description = re.ReplaceAllString(item.Description, "")
	// }
	// fmt.Println(feed)

	var feed Feed
	xml.Unmarshal(data, &feed)

	rsp, err := json.Marshal(feed)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rsp)

	// return
}

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
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

	return data, nil
}

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Plant struct {
// 	XMLName xml.Name `xml:"plant"`
// 	Id      int      `xml:"id,attr"`
// 	Name    string   `xml:"name"`
// 	Origin  []string `xml:"origin"`
// }

// func (p Plant) String() string {
// 	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
// 		p.Id, p.Name, p.Origin)
// }

// func main() {
// 	coffee := &Plant{Id: 27, Name: "Coffee"}
// 	coffee.Origin = []string{"Ethiopia", "Brazil"}

// 	out, _ := xml.MarshalIndent(coffee, " ", "  ")
// 	fmt.Println(string(out))
// 	fmt.Println("====================1=================")
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Println(xml.Header + string(out))
// 	fmt.Println("====================2=================")
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()

// 	var p Plant
// 	if err := xml.Unmarshal(out, &p); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(p)
// 	fmt.Println("====================3=================")
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()

// 	tomato := &Plant{Id: 81, Name: "Tomato"}
// 	tomato.Origin = []string{"Mexico", "California"}

// 	type Nesting struct {
// 		XMLName xml.Name `xml:"nesting"`
// 		Plants  []*Plant `xml:"parent>child>plant"`
// 	}

// 	nesting := &Nesting{}
// 	nesting.Plants = []*Plant{coffee, tomato}

// 	out, _ = xml.MarshalIndent(nesting, " ", "  ")
// 	fmt.Println(string(out))
// 	fmt.Println("====================4=================")
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()
// }
