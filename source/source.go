package source

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Description struct {
	location string
	format   string
	distance string
}

func Reader(d Description) (io.ReadCloser, error) {
	if d.distance == "local" {
		return os.Open(d.location)
	}
	if d.distance == "remote" {
		resp, err := http.Get(d.location)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	return nil, errors.New(fmt.Sprintf(
		"Unknown source distance `%s`",
		d.distance,
	))
}
