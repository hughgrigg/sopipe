package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func PostsFromLocation(location string) ([]Post, error) {
	var posts []Post
	r, err := locationToReader(location)
	defer r.Close()
	if err != nil {
		return posts, err
	}

	for _, sourcePoster := range sourcePosters {
		posts, err = sourcePoster(r)
		if err == nil {
			return posts, nil
		}
	}

	return posts, fmt.Errorf("Failed to get posts from location `%s`", location)
}

func PostToTabbed(p Post) string {
	return fmt.Sprintf("%s\t%s", p.Body, p.Links[0].String())
}

type Post struct {
	Body     string
	Links    []url.URL
	Mentions []Mention
	HashTags []HashTag
}

type Image struct {
	src url.URL
}

type HashTag struct {
	string
}

type Mention struct {
	string
}

var sourcePosters []sourcePoster = []sourcePoster{
	RSSToPosts,
}

type sourcePoster func(r io.Reader) ([]Post, error)

type postFormatter func(p Post) string

func locationToReader(location string) (io.ReadCloser, error) {
	if parseUrl, err := url.Parse(location); err == nil && parseUrl.IsAbs() {
		get, err := http.Get(parseUrl.String())
		return get.Body, err
	}
	if openFile, err := os.Open(location); err == nil {
		return openFile, nil
	}
	return nil, fmt.Errorf("Failed to open location %s", location)
}
