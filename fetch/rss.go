package fetch

import (
	"encoding/xml"
	"io"
	"net/url"
	"strings"
	"time"
)

func RSSToPosts(r io.Reader) ([]Post, error) {
	var posts []Post
	rss, err := ParseRSS(r)
	if err != nil {
		return posts, err
	}
	for _, item := range rss.Channel.Items {
		posts = append(posts, Post{
			Body:  strings.TrimSpace(item.Title),
			Links: []url.URL{item.Link.URL},
		})
	}
	return posts, nil
}

type RSSFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	XMLName xml.Name  `xml:"channel"`
	Title   string    `xml:"title"`
	Items   []RSSItem `xml:"item"`
}

type RSSItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        URL      `xml:"link"`
	PubDate     RSSDate  `xml:"pubDate"`
	Description string   `xml:"description"`
}

type URL struct {
	url.URL
}

func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := url.Parse(strings.TrimSpace(v))
	if err != nil {
		return err
	}
	*u = URL{*parsed}
	return nil
}

type RSSDate struct {
	time.Time
}

func (rd *RSSDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := time.Parse(time.RFC1123Z, v)
	if err != nil {
		return err
	}
	*rd = RSSDate{parsed}
	return nil
}

func ParseRSS(r io.Reader) (RSSFeed, error) {
	var rssFeed RSSFeed
	err := xml.NewDecoder(r).Decode(&rssFeed)
	return rssFeed, err
}
