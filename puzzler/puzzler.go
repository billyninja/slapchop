package puzzler

import (
	"encoding/json"
	"flag"
	"github.com/billyninja/slapchop/chopper"
	"github.com/go-resty/resty"
	"net/http"
)

var FlagPuzzlerHost = flag.String("puzzler", "localhost:8000", "Puzzler Service remote url")

func CreatePuzzle(username string, tiles []*chopper.TileEntry) (int, *http.Response, error) {
	pieces, _ := json.Marshal(tiles)

	resp, err := resty.R().
		SetBasicAuth("admin", "teste123").
		SetFormData(map[string]string{
		"username": username,
		"pieces":   string(pieces),
	}).
		SetHeader("Content-Type", "application/json").
		Post("http://" + *FlagPuzzlerHost + "/puzzles/")
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.RawResponse, err
}
