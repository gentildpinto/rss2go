package main

import "fmt"

func main() {
	feed, _ := Rss2Go("https://g1.globo.com/rss/g1")
	println("Hi!")
	fmt.Println(feed)
}
