package puzzler

import (
	"github.com/billyninja/slapchop/actions"
	"github.com/go-resty/resty"
)

func CreatePuzzle(actions.CreateResponse) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		Post("http://localhost:8000")

	println(resp, err)

	return
}
