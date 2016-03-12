package fetch

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPostsFromRSSURL(t *testing.T) {
	rssFile, err := os.Open("./test/rss.xml")
	defer rssFile.Close()
	if err != nil {
		t.Errorf("Failed to open rss data file: %s", err)
	}
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			io.Copy(w, rssFile)
		},
	))
	defer ts.Close()

	posts, err := PostsFromLocation(ts.URL)
	if err != nil {
		t.Errorf("Failed to get posts from RSS URL: %s", err)
	}
	assert.Len(t, posts, 2)
}

func TestFileToReader(t *testing.T) {
	reader, err := locationToReader("./test/rss.xml")
	defer reader.Close()
	if err != nil {
		t.Errorf("Failed to get reader for file: %s", err)
	}
	br := bufio.NewReader(reader)
	firstLine, err := br.ReadString('\n')
	if err != nil {
		t.Errorf("Failed to get first line from file reader: %s", err)
	}
	assert.Equal(
		t,
		"<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\" ?>\n",
		firstLine,
	)
}

func TestURLToReader(t *testing.T) {
	sourceContent := "foo post source"
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, sourceContent)
		},
	))
	defer ts.Close()

	reader, err := locationToReader(ts.URL)
	if err != nil {
		t.Errorf("Failed to get reader for URL: %s", err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	retrievedContent := buf.String()
	if err != nil {
		t.Errorf("Failed to get content of URL reader: %s", err)
	}
	assert.Equal(t, sourceContent, retrievedContent)
}
