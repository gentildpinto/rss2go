# rss2go

`rss2go` is a Go package designed to fetch and parse RSS feeds. It provides a simple API to request RSS data from a URL and unmarshals the XML content into Go structs.

## Features

- Fetches RSS feeds via HTTP requests
- Parses and unmarshals RSS feed data from XML into Go structs
- Supports media content in RSS feed items
- Handles common RSS formats including Atom links
- Cleans up CDATA tags from the response

## Installation

To use `rss2go`, you need to have Go installed. You can install the package with the following command:

```bash
go get github.com/gentildpinto/rss2go
```

## Usage

Import the package in your Go project and call the Rss2Go function to fetch and parse an RSS feed:

```go
package main

import (
    "fmt"
    "log"

    "github.com/gentildpinto/rss2go"
)

func main() {
    feedUrl := "https://example.com/rss"
    feed, err := rss2go.Rss2Go(feedUrl)
    if err != nil {
        log.Fatalf("Error fetching feed: %v", err)
    }

    fmt.Println("Feed Title:", feed.Channel.Title)
    for _, item := range feed.Channel.Items {
        fmt.Printf("Title: %s\nLink: %s\n", item.Title, item.Link)
    }
}
```

## Structs

The following Go structs are used to represent the structure of an RSS feed:

- **Feed**: The root RSS element.
- **Channel**: Contains metadata about the RSS feed.
- **Item**: Represents each entry in the feed.
- **Image**: Optional image metadata for the feed.
- **MediaContent**: Represents media content (if any) in an item.
- **AtomLink**: Used for Atom feed compatibility.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you encounter any bugs or have suggestions for new features.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
