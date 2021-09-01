package main

import (
	"encoding/xml"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func findInDoc(topic, doc string) int {
	var found int
	items, err := read(doc)
	if err != nil {
		return 0
	}

	for _, item := range items {
		if strings.Contains(item.Description, topic) {
			found++
		}
	}
	return found
}

// find is a sequential find
func find(topic string, docs []string) int {
	var found int
	for _, doc := range docs {
		found += findInDoc(topic, doc)
	}
	return found
}

// findConcurrent is the concurrent implementation of find
// it uses a pooling pattern where a channel is used to feed a pool of
// goroutins
func findConcurrent(goroutines int, topic string, docs []string) int {
	var found int64

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			var lFound int
			for doc := range ch {
				findInDoc(topic, doc)
			}
			atomic.AddInt64(&found, int64(lFound))
			wg.Done()
		}()
	}
	wg.Wait()
	return int(found)
}

func generateList(totalDocs int) []string {
	docs := make([]string, totalDocs)
	for i := 0; i < totalDocs; i++ {
		docs[i] = "test.xml"
	}
	return docs
}

func read(doc string) ([]item, error) {
	time.Sleep(time.Millisecond)
	var d document

	if err := xml.Unmarshal([]byte(file), &d); err != nil {
		return nil, err
	}

	return d.Channel.Items, nil
}

var file = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an amazing language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []item   `xml:"item"`
}

type document struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
	URI     string
}
