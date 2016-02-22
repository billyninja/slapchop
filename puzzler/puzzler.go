package puzzler

import (
	"encoding/json"
	"github.com/billyninja/slapchop/chopper"
	"github.com/go-resty/resty"
)

func CreatePuzzle(username string, tiles []*chopper.TileEntry) {

	pieces, _ := json.Marshal(tiles)

	resp, err := resty.R().
		SetBasicAuth("admin", "teste123").
		SetFormData(map[string]string{
		"username": username,
		"pieces":   string(pieces),
	}).
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:8000/solutions/")

	return
}
