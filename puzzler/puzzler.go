package puzzler

import (
	"github.com/go-resty/resty"
	"github.com/slapchop/actions"
)


func CreatePuzzle(actions.CreateResponse) {
	resp, err := resty.R().
    	SetHeader("Content-Type", "application/json").
    	SetBody(`{"username":"testuser", "password":"testpass"}`).
      	Post("http://localhost:8000")

	return
}
