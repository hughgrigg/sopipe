package source

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

func TestFetchesRemoteSource(t *testing.T) {
	content := []byte("remote source content")
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(content)
		},
	))
	defer ts.Close()
	desc := Description{
		location: ts.URL,
		distance: "remote",
		format:   "rss",
	}
	reader, err := Reader(desc)
	readContent, err := ioutil.ReadAll(reader)
	check(t, err)
	defer reader.Close()
	assert.Equal(t, content, readContent)
}

func TestFetchesLocalSource(t *testing.T) {
	location := "test_source.txt"
	content := []byte("local source content")
	f, err := os.Create(location)
	check(t, err)
	defer os.Remove(location)
	_, err = f.Write(content)
	check(t, err)
	err = f.Close()
	check(t, err)
	desc := Description{
		location: location,
		distance: "local",
		format:   "rss",
	}
	reader, err := Reader(desc)
	check(t, err)
	readContent, err := ioutil.ReadAll(reader)
	check(t, err)
	defer reader.Close()
	assert.Equal(t, content, readContent)
}
