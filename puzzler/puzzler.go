package puzzler

import (
	"github.com/billyninja/slapchop/chopper"
	"github.com/go-resty/resty"
)

func CreatePuzzle(chopper.CreateResponse) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"admin", "password":"testpass"}`).
		Post("http://localhost:8000/solutions/")

	println(resp, err)

	return
}
