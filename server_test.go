package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/billyninja/slapchop/actions"
	"github.com/billyninja/slapchop/chopper"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var testPort = "3002"
var testHost = "http://localhost:" + testPort
var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func TestChain(t *testing.T) {

	server_control := InitServer(testPort)

	// CREATE
	req1, err := actions.NewfileUploadRequest(testHost+"/chopit/testuser", "uploadfile", "test/gopher.jpg")
	resp1, err := client.Do(req1)
	if err != nil {
		t.Error(err)
		return
	}
	body1, _ := ioutil.ReadAll(resp1.Body)
	cr := chopper.CreateResponse{}
	err = json.Unmarshal(body1, &cr)
	if err != nil {
		t.Error(err)
		return
	}
	resp1.Body.Close()

	// READ-ONE
	req2, err := http.NewRequest("GET", testHost+"/chopit/testuser/"+cr.ChopId, nil)
	resp2, err := client.Do(req2)
	if err != nil {
		t.Error(err)
		return
	}
	body2, _ := ioutil.ReadAll(resp2.Body)
	ror := chopper.ReadAllResponse{}
	err = json.Unmarshal(body2, &ror)
	if err != nil {
		t.Error(err)
		return
	}
	resp2.Body.Close()

	// DELETE
	req3, err := http.NewRequest("DELETE", testHost+"/chopit/testuser/"+cr.ChopId, nil)
	resp3, err := client.Do(req3)
	if err != nil {
		t.Error(err)
		return
	}
	body3, _ := ioutil.ReadAll(resp3.Body)
	println(">>>", string(body3))
	resp3.Body.Close()

	// Closing the server
	server_control <- os.Interrupt
}
