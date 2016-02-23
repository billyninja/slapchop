package puzzler

import (
	"encoding/json"
	"github.com/billyninja/slapchop/chopper"
	"github.com/go-resty/resty"
	"net/http"
)

func CreatePuzzle(username string, tiles []*chopper.TileEntry) (int, *http.Response, error) {

	pieces, _ := json.Marshal(tiles)

	resp, err := resty.R().
		SetBasicAuth("admin", "teste123").
		SetFormData(map[string]string{
		"username": username,
		"pieces":   string(pieces),
	}).
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:8000/puzzles/")
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.RawResponse, err
}
