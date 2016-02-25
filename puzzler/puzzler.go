package puzzler

import (
	"encoding/json"
	"github.com/billyninja/slapchop/chopper"
	"github.com/go-resty/resty"
	"strings"
)

type CreateResponse struct {
	PuzzleHref   string `json:"puzzle_href"`
	SolutionHref string `json:"solution_href"`
}

func CreatePuzzle(puzzler_host string, username string, tiles []*chopper.TileEntry) (int, []byte, error) {
	pieces, _ := json.Marshal(tiles)

	auth_and_host := strings.Split(puzzler_host, "@")
	println(auth_and_host[0], auth_and_host[1])
	auth := strings.Split(auth_and_host[0], ":")

	resp, err := resty.R().
		SetBasicAuth(auth[0], auth[1]).
		SetFormData(map[string]string{
		"username": username,
		"pieces":   string(pieces),
	}).
		SetHeader("Content-Type", "application/json").
		Post("http://" + auth_and_host[1] + "/puzzles/")
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), err
}
