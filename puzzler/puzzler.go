package puzzler

import (
	"github.com/billyninja/slapchop/actions"
	"github.com/go-resty/resty"
)

func CreatePuzzle(actions.CreateResponse) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"admin", "password":"testpass"}`).
		Post("http://localhost:8000/solutions/")

	println(resp, err)

	return
}
