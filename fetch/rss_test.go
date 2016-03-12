package fetch

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var expectedRSSItems = []map[string]string{
	{
		"title": "Refer to other memories",
		"link":  "https://www.chineseboost.com/blog/refer-to-other-memories/",
	},
	{
		"title": "5 ways to optimise your Chinese flashcards",
		"link":  "https://www.chineseboost.com/blog/optimise-chinese-flashcards/",
	},
}

func TestParseRSS(t *testing.T) {
	rssData, err := os.Open("./test/rss.xml")
	if err != nil {
		t.Error("Failed to read RSS test data")
	}
	rssFeed, err := ParseRSS(rssData)
	if err != nil {
		t.Errorf("Failed to parse RSS into RSSFeed: %s", err)
	}
	assert.Equal(t, "Chinese Boost", rssFeed.Channel.Title)
	assert.Equal(t, 2, len(rssFeed.Channel.Items))
	for i, item := range rssFeed.Channel.Items {
		assert.Equal(t, expectedRSSItems[i]["title"], item.Title)
		assert.Equal(t, expectedRSSItems[i]["link"], item.Link.String())
	}
}

func TestRSSToPosts(t *testing.T) {
	rssData, err := os.Open("./test/rss.xml")
	if err != nil {
		t.Error("Failed to read RSS test data")
	}
	posts, err := RSSToPosts(rssData)
	if err != nil {
		t.Errorf("Failed to make posts from RSSFeed: %s", err)
	}
	assert.Equal(t, 2, len(posts))
	for i, post := range posts {
		assert.Equal(t, expectedRSSItems[i]["title"], post.Body)
		assert.Equal(t, expectedRSSItems[i]["link"], post.Links[0].String())
	}
}

func TestNotRSS(t *testing.T) {
	notRSS := strings.NewReader(
		`<?xml version="1.0" encoding="utf-8" ?>
		<not_rss></not_rss>`,
	)
	posts, err := RSSToPosts(notRSS)
	assert.Empty(t, posts)
	assert.NotEmpty(t, err)
	assert.Equal(
		t,
		"expected element type <rss> but have <not_rss>",
		err.Error(),
	)
}
